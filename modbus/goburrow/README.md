# check device

```
$ ./check_dev.sh
syspath   : /sys/bus/usb/devices/usb3/3-1/3-1:1.0/ttyUSB0/tty/ttyUSB0
ID_SERIAL: 1a86_USB_Serial
1a86:7523 /dev/ttyUSB0
Bus 003 Device 002: ID 1a86:7523 QinHeng Electronics HL-340 USB-Serial adapter
```

# run

```
go build

sudo ./goburrow --action=get

sudo ./goburrow --action=set --port=1
```