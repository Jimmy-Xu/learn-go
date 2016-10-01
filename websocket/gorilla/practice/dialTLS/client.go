// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.


package main

import (
	"io/ioutil"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"
	"github.com/gorilla/websocket"
	SignUtil "github.com/Jimmy-Xu/learn-go/websocket/gorilla/practice/dialTLS/util"
)

var addr = flag.String("addr", "localhost:8888", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: *addr, Path: "/echo"}
	log.Printf("connecting to %s", u.String())

	//load ca
	pool := x509.NewCertPool()
	caCertPath := "ssl/ca.crt"
	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	//load client cert
	cliCrt, err := tls.LoadX509KeyPair("ssl/client.crt", "ssl/client.key")
	if err != nil {
		fmt.Println("Loadx509keypair err:", err)
		return
	}

	config := &tls.Config{
		RootCAs:      pool,
		Certificates: []tls.Certificate{cliCrt},
	}
	dialer := websocket.Dialer{
		TLSClientConfig: config,
	}
	
	fmt.Printf("URL:%v\n", u.String())
	req, err := http.NewRequest("GET", u.String(), nil)
	req.URL = &u

	accessKey := "6DVNAWRWDP6NUVGEOLKGJ9YV"
	secretKey := "48K101y5gAPce7rZVPlWNiOOjk7BA5kUPao7qeQQ"
	req = SignUtil.Sign4(accessKey, secretKey, req)

	c, _, err := dialer.Dial(u.String(), req.Header)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer c.Close()
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")
		// To cleanly close a connection, a client should send a close
		// frame and wait for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
				select {
				case <-done:
				case <-time.After(time.Second):
				}
			c.Close()
			return
		}
	}
}
