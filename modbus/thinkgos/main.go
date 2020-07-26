package main

import (
	"encoding/json"
	"flag"

	"github.com/goburrow/serial"
	"github.com/sirupsen/logrus"
	modbus "github.com/thinkgos/gomodbus"
)

var (
	slaveID byte = 1
	//address  uint16 = 1
	//quantity uint16 = 1
)

func main() {
	action := flag.String("action", "get", "get or set")
	address := flag.Uint("address", 1, "address")
	quantity := flag.Uint("quantity", 1, "quantity")
	port := flag.Int("port", 1, "port number, 0 or 1")
	debug := flag.Bool("debug", false, "enable debug")
	flag.Parse()

	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	p := modbus.NewRTUClientProvider(modbus.WithEnableLogger(),
		modbus.WithSerialConfig(serial.Config{
			//Address:  "/dev/ttyUSB0",
			Address:  "COM3",
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
		getPort(client, uint16(*address), uint16(*quantity))
	case "set":
		setPort(client, uint16(*address), uint16(*port))
	default:
		logrus.Errorf("%s is unknown action", *action)
	}
}

func getPort(client modbus.Client, address uint16, quantity uint16) {
	value, err := client.ReadHoldingRegisters(slaveID, address, quantity)
	if err != nil {
		logrus.WithError(err).Error("failed to read holding registers")
	} else {
		logrus.Infof("%#v\n", value)
	}
}

func setPort(client modbus.Client, address uint16, port uint16) {
	err := client.WriteSingleRegister(slaveID, address, port)
	if err != nil {
		logrus.WithError(err).Error("failed to write single register")
	} else {
		logrus.Infof("write single register OK")
	}
}
