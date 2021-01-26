package main

import (
	"io"
	"net"
)

func main() {
	// '@' indicates the socket held in an abstract namespace
	// which doesn't belong to a file in the filesystem
	abstractUnixSocket := "@criu.sock"

	ln, err := net.Listen("unix", abstractUnixSocket)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		// handle conn
		go func(conn net.Conn) {
			defer conn.Close()
			io.Copy(conn, conn)
		}(conn)
	}
}
