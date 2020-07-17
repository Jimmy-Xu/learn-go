package main

import (
	"encoding/json"
	"flag"

	"github.com/goburrow/serial"
	"github.com/sirupsen/logrus"
	modbus "github.com/thinkgos/gomodbus"
)

var (
	slaveID  byte   = 1
	address  uint16 = 1
	quantity uint16 = 1
)

func main() {
	action := flag.String("action", "get", "get or set")
	port := flag.Int("port", 1, "port number, 1-4")
	debug := flag.Bool("debug", false, "enable debug")
	flag.Parse()

	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	p := modbus.NewRTUClientProvider(modbus.WithEnableLogger(),
		modbus.WithSerialConfig(serial.Config{
			Address:  "/dev/ttyUSB0",
			BaudRate: 9600,
			DataBits: 8,
			StopBits: 1,
			Parity:   "N",
			Timeout:  modbus.SerialDefaultTimeout,
		}))

	buf, _ := json.Marshal(p.Config)

	logrus.Debugf("config: %v\n", string(buf))
	client := modbus.NewClient(p)
	err := client.Connect()
	if err != nil {
		logrus.WithError(err).Error("modbus connect failed")
		return
	}
	defer client.Close()
	if *debug {
		client.LogMode(true)
	}

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
	value, err := client.ReadHoldingRegisters(slaveID, address, quantity)
	if err != nil {
		logrus.WithError(err).Error("failed to read holding registers")
	} else {
		logrus.Infof("%#v\n", value)
	}
}

func setPort(client modbus.Client, port int) {
	err := client.WriteSingleRegister(slaveID, address, uint16(port))
	if err != nil {
		logrus.WithError(err).Error("failed to write single register")
	} else {
		logrus.Infof("write single register OK")
	}
}
