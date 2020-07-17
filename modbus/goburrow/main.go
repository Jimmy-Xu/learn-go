package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/goburrow/modbus"
	"github.com/sirupsen/logrus"
)

// 定义了EnvMap，其目的是存放各个配置项，便于程序随时使用，下面紧跟的是每个具体的配置项的key
var EnvMap = make(map[string]string)
var RtuDevice = "rtudevice"
var BaudRate = "baudrate"
var DataBits = "databits"
var Parity = "parity"
var StopBits = "stopbits"
var SlaveId = "slaveid"
var SerialTimeout = "serialtimeout"
var Address = "address"

// 定义上面每个key的初始化值
const (
	DefaultRtuDevcie     = "/dev/ttyUSB0"
	DefaultBaudRate      = "19200"
	DefaultDataBits      = "8"
	DefaultParity        = "N"
	DefaultStopBits      = "1"
	DefaultSlaveId       = "1"
	DefaultSerialTimeout = "5"
	DefaultAddress       = "1"
)

// 定义配置文件的名称
const (
	workerConf = "modbus.conf"
)

const (
	address  = 1
	quantity = 1
)

func main() {
	action := flag.String("action", "get", "get or set")
	port := flag.Int("port", 1, "port number, 1-4")
	debug := flag.Bool("debug", false, "enable debug")
	flag.Parse()

	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	// 读参数
	initConfMap()

	// 根据串口设备初始化一个handler
	handler := modbus.NewRTUClientHandler(EnvMap[RtuDevice])

	// 配置handler的波特率
	baudRate, _ := strconv.Atoi(EnvMap[BaudRate])
	handler.BaudRate = baudRate
	// 数据位
	dataBits, _ := strconv.Atoi(EnvMap[DataBits])
	handler.DataBits = dataBits

	// // 校验位
	// handler.Parity = EnvMap[Parity]
	handler.Parity = "N"

	// 停止位
	stopBits, _ := strconv.Atoi(EnvMap[StopBits])
	handler.StopBits = stopBits

	// modbus SlaveId
	slaveId, _ := strconv.Atoi(EnvMap[SlaveId])
	handler.SlaveId = byte(slaveId)

	// 超时时间
	serialTimeout, _ := strconv.Atoi(EnvMap[SerialTimeout])
	handler.Timeout = time.Duration(serialTimeout) * time.Second

	err := handler.Connect()
	if err != nil {
		logrus.WithError(err).Errorf("failed to connect %v", EnvMap[RtuDevice])
		os.Exit(1)
	}
	defer handler.Close()

	client := modbus.NewClient(handler)

	switch *action {
	case "get":
		getPort(client)
	case "set":
		setPort(client, *port)
	default:
		logrus.Errorf("%s is unknown action", *action)
	}

}

func getPort(client modbus.Client) {
	value, err := client.ReadHoldingRegisters(address, quantity)
	if err != nil {
		logrus.WithError(err).Error("failed to read holding registers")
	} else {
		logrus.Infof("%#v\n", value)
	}
}

func setPort(client modbus.Client, port int) {
	results, err := client.WriteSingleRegister(address, uint16(port))
	if err != nil {
		logrus.WithError(err).Error("failed to write single register")
	} else {
		logrus.Infof("write single register OK, result:%v", results)
	}
}

func initConfMap() {

	buf, err := ioutil.ReadFile(workerConf)
	if err != nil {
		logrus.WithError(err).Fatal("failed to read config file")
	}
	strs := string(buf)
	if strs != "" {
		result := []string{}
		for _, lineStr := range strings.Split(strs, "\n") {
			lineStr = strings.TrimSpace(lineStr)
			if lineStr == "" {
				continue
			}
			result = strings.Split(lineStr, "=")
			k := result[0]
			v := result[1]
			EnvMap[k] = v
		}
	}
	// give a default value
	if EnvMap[RtuDevice] == "" {
		EnvMap[RtuDevice] = DefaultRtuDevcie
	}
	if EnvMap[BaudRate] == "" {
		EnvMap[BaudRate] = DefaultBaudRate
	}
	if EnvMap[DataBits] == "" {
		EnvMap[DataBits] = DefaultDataBits
	}
	if EnvMap[Parity] == "" {
		EnvMap[Parity] = DefaultParity
	}
	if EnvMap[StopBits] == "" {
		EnvMap[StopBits] = DefaultStopBits
	}
	if EnvMap[SlaveId] == "" {
		EnvMap[SlaveId] = DefaultSlaveId
	}
	if EnvMap[SerialTimeout] == "" {
		EnvMap[SerialTimeout] = DefaultSerialTimeout
	}
	if EnvMap[Address] == "" {
		EnvMap[Address] = DefaultAddress
	}

	logrus.Debugf("EnvMap:%v", EnvMap)
}

func bytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
