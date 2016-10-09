package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	SignUtil "github.com/Jimmy-Xu/learn-go/websocket/gorilla/practice/hyper/util"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"
)


var addr = flag.String("addr", "147.75.195.37:6443", "http service address")
var accessKey = flag.String("accessKey", "", "hyper access key")
var secretKey = flag.String("secretKey", "", "hyper secret key")
var u = url.URL{Scheme: "wss", Host: *addr, Path: "/events/ws"}

func main() {
	flag.Parse()
	log.SetFlags(0)

	//check accessKey and secretKey
	if *accessKey == "" {
		log.Printf("accessKey can not be empty!")
		return
	}
	if *secretKey == "" {
		log.Printf("secretKey can not be empty!")
		return
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	log.Printf("connecting to %s", u.String())

	//add sign to header
	req, err := http.NewRequest("GET", u.String(), nil)
	req.URL = &u
	req = SignUtil.Sign4(*accessKey, *secretKey, req)

	//connect to websocket server
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	dialer := websocket.Dialer{
		TLSClientConfig: config,
	}
	c, _, err := dialer.Dial(u.String(), req.Header)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	done := make(chan struct{})

	//process websocket message
	go func() {
		defer c.Close()
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			//way1: show raw result
			//log.Printf("recv: %s", message)

			//way2: show pretty result
			var dat map[string]interface{}
			if err := json.Unmarshal([]byte(message), &dat); err != nil {
				panic(err)
			}
			b, err := json.MarshalIndent(dat, "", "  ")
			if err != nil {
				panic(err)
			}
			log.Printf("recv[json]: %v\n\n", string(b[:]))
		}
	}()

	for {
		select {
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
