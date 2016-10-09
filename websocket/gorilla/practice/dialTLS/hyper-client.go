package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	SignUtil "github.com/Jimmy-Xu/learn-go/websocket/gorilla/practice/dialTLS/util"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"
)

var addr = flag.String("addr", "147.75.195.37:6443", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: *addr, Path: "/events/ws"}
	log.Printf("connecting to %s", u.String())

	config := &tls.Config{
		InsecureSkipVerify: true,
	}

	dialer := websocket.Dialer{
		TLSClientConfig: config,
	}

	//fmt.Printf("URL:%v\n", u.String())
	req, err := http.NewRequest("GET", u.String(), nil)
	req.URL = &u

	accessKey := "6DVNAWRWDP6NUVGEOLKGJ9YV"
	secretKey := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

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
			//show raw result
			//log.Printf("recv: %s", message)

			//show pretty result
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

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		// case t := <-ticker.C:
		// 	err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
		// 	if err != nil {
		// 		log.Println("write:", err)
		// 		return
		// 	}
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
