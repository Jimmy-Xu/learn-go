package main

import (
	"fmt"
	"net"
)

type UnixSocket struct {
	filename string
	bufsize  int
	handler  func(string) string
}

func main() {
	 //声明unixsocket
	 us := NewUnixSocket("/tmp/us.socket")
	 //发送数据unixsocket并返回服务端处理结果
	 r := us.ClientSendContext("hello")
	 fmt.Println("response:" + r)
}

func NewUnixSocket(filename string, size ...int) *UnixSocket {
	size1 := 10480
	if size != nil {
		size1 = size[0]
	}
	us := UnixSocket{filename: filename, bufsize: size1}
	return &us
}

func (this *UnixSocket) ClientSendContext(context string) string {
	addr, err := net.ResolveUnixAddr("unix", this.filename)
	if err != nil {
		panic("Cannot resolve unix addr: " + err.Error())
	}
	//拔号
	c, err := net.DialUnix("unix", nil, addr)
	if err != nil {
		panic("DialUnix failed.")
	}
	//写出
	_, err = c.Write([]byte(context))
	if err != nil {
		panic("Writes failed.")
	}
	//读结果
	buf := make([]byte, this.bufsize)
	nr, err := c.Read(buf)
	if err != nil {
		panic("Read: " + err.Error())
	}
	return string(buf[0:nr])
}