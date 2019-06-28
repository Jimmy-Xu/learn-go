package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("unix", "/var/run/docker.sock")

	if err != nil {
		panic(err)
	}
	fmt.Fprintf(conn, "GET /images/json HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\t')
	if err != nil {
		log.Println(err)
	}

	fmt.Print("Message from server: "+status)
}