package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/mdlayher/vsock"
	"github.com/sirupsen/logrus"
)

const (
	VSockSocketScheme = "vsock"
	vsockPortBegin    = 40000
	vsockPortEnd      = 49999
)

type vsockProxy struct {
	host string
	port int

	vsockURL  string
	sandboxID string

	isRunning bool
	quit      chan bool
	mutex     sync.Mutex

	listener net.Listener
}

func main() {
	logrus.Infof("vsock: %v", os.Args[1])
	proxy := newVSOCKProxy(os.Args[1])
	proxy.StartProxy()
}

func newVSOCKProxy(vsockURL string) *vsockProxy {
	return &vsockProxy{
		host:     "0.0.0.0",
		quit:     make(chan bool),
		vsockURL: vsockURL,
	}
}

func (v *vsockProxy) StartProxy() error {
	var err error
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		logrus.Infof("startTCPServer")
		err = v.startTCPServer()
		if err != nil {
			logrus.Infof("failed to startTCPServer, error:%v", err)
		}
		wg.Done()
	}()
	wg.Wait()

	return err
}

func (v *vsockProxy) startTCPServer() error {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	var err error
	var found = false

	//随机搜索一个可用的tcp port
	lenPort := vsockPortEnd - vsockPortBegin + 1
	offset := rand.Intn(lenPort)
	for i := vsockPortBegin; i <= vsockPortEnd; i++ {
		v.port = i + offset
		if v.port > vsockPortEnd {
			v.port -= lenPort
		}
		logrus.Infof("detect free port: %v", v.port)
		v.listener, err = net.Listen("tcp4", fmt.Sprintf("%s:%d", v.host, v.port))
		if err != nil {
			if strings.Contains(err.Error(), "address already in use") {
				logrus.Errorf("port %v in use, retry", v.port)
				continue
			} else {
				logrus.Errorf("Failed to start TCP server on port %v", v.port)
				return err
			}
		} else {
			found = true
			break
		}
	}
	if !found {
		return errors.New("No available port for VSOCK proxy")
	}

	v.isRunning = true
	logrus.Infof("VSOCK(%v) proxy started: port %v", v.vsockURL, v.port)

	v.serve()

	return nil
}

func (v *vsockProxy) serve() {
	for {
		select {
		case <-v.quit:
			logrus.Infof("Server quit.")
			return
		default:
			logrus.Infof("Waiting for VSOCK connection...")
			tc, err := v.listener.Accept()
			if err != nil {
				logrus.Infof("Listener accept failed")
				continue
			}
			go v.connectVSock(tc)
		}
	}
}

func (v *vsockProxy) connectVSock(tc net.Conn) {
	//连接vsock
	cid, port, err := parseGrpcVsockAddr(v.vsockURL)
	uc, err := vsock.Dial(cid, port)
	if err != nil {
		logrus.Infof("failed to connect to vsock: %v (cid:%v port:%v), error:%v", v.vsockURL, cid, port, err)
		return
	}
	//转发
	go v.ioCopy(tc, uc)
	go v.ioCopy(uc, tc)
	logrus.Infof("io copy started")
}

func (v *vsockProxy) ioCopy(dst, src net.Conn) {
	defer dst.Close()
	written, err := io.Copy(dst, src)
	if err != nil {
		logrus.Infof("io copy failed")
		return
	}
	logrus.Infof("io copy written: %v", written)
}

func (v *vsockProxy) StopProxy() {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	if !v.isRunning {
		logrus.Info("VSOCK proxy is not running. Do nothing.")
		return
	}

	logrus.Info("Stopping VSOCK proxy")
	close(v.quit)
	v.listener.Close()
	v.port = 0
	logrus.Info("VSOCK proxy stopped.")
}

func parseGrpcVsockAddr(sock string) (uint32, uint32, error) {
	sp := strings.Split(sock, ":")
	sp[1] = strings.Replace(sp[1], "//", "", 1)
	logrus.Infof("parse vsock url:%v, sp:%v", sock, sp)
	if len(sp) != 3 {
		return 0, 0, fmt.Errorf("invalid vsock address: %s", sock)
	}
	if sp[0] != VSockSocketScheme {
		return 0, 0, fmt.Errorf("invalid vsock URL scheme: %s", sp[0])
	}

	cid, err := strconv.ParseUint(sp[1], 10, 32)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid vsock cid: %s", sp[1])
	}
	port, err := strconv.ParseUint(sp[2], 10, 32)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid vsock port: %s", sp[2])
	}

	return uint32(cid), uint32(port), nil
}
