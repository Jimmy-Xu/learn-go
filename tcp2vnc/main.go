package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

const (
	vncPortBegin        = 59000
	vncPortEnd          = 59999
)

type vncProxy struct {
	host      string
	port      int
	unixSock  string
	sandboxID string

	isRunning bool
	quit      chan bool
	mutex     sync.Mutex

	listener net.Listener
}

func main() {
	v := newVNCProxy(os.Args[1])
	v.StartProxy()
}

func newVNCProxy(sock string) *vncProxy {
	return &vncProxy{
		host:      "0.0.0.0",
		quit:      make(chan bool),
		unixSock:  sock,
	}
}


func (v *vncProxy) StartProxy() error {
	var err error
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		err = v.startTCPServer()
		wg.Done()
	}()
	wg.Wait()

	return err
}

func (v *vncProxy) startTCPServer() error {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	var err error
	var found = false

	lenPort := vncPortEnd - vncPortBegin + 1
	offset := rand.Intn(lenPort)
	for i := vncPortBegin; i <= vncPortEnd; i++ {
		v.port = i + offset
		if v.port > vncPortEnd {
			v.port -= lenPort
		}
		v.listener, err = net.Listen("tcp4", fmt.Sprintf("%s:%d", v.host, v.port))
		if err != nil {
			if strings.Contains(err.Error(), "address already in use") {
				logrus.Debugf("port %v in use, retry", v.port)
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
		return errors.New("No available port for VNC proxy")
	}

	go v.serve()
	v.isRunning = true
	logrus.Debugf("VNC proxy started: port %v", v.port)

	return nil
}

func (v *vncProxy) serve() {
	for {
		select {
		case <-v.quit:
			logrus.Debug("Server quit.")
			return
		default:
			logrus.Debug("Waiting for VNC connection...")
			tc, err := v.listener.Accept()
			if err != nil {
				logrus.Infof("Listener accept failed")
				continue
			}
			go v.connectSock(tc)
		}
	}
}

func (v *vncProxy) connectSock(tc net.Conn) {
	uc, err := net.Dial("unix", v.unixSock)
	if err != nil {
		logrus.Errorf("failed to connect to unix sock: %v", v.unixSock)
		return
	}
	go v.ioCopy(tc, uc)
	go v.ioCopy(uc, tc)
	logrus.Debugf("io copy started")
}

func (v *vncProxy) ioCopy(dst, src net.Conn) {
	defer dst.Close()
	written, err := io.Copy(dst, src)
	if err != nil {
		logrus.Debugf("io copy failed")
		return
	}
	logrus.Debugf("io copy written: %v", written)
}

func (v *vncProxy) StopProxy() {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	if !v.isRunning {
		logrus.Info("VNC proxy is not running. Do nothing.")
		return
	}

	logrus.Info("Stopping VNC proxy")
	close(v.quit)
	v.listener.Close()
	v.port = 0
	logrus.Infof("VNC proxy stopped.")
}
