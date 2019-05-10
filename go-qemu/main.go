package main

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/jimmy-xu/learn-go/go-qemu/qmp"
)

func main() {
	socket:="/run/vc/vm/19cc5ad191371184541dcc0c0ff197a3a59247cd76025a511f25aa6cc59bd897/qga.sock"
	if len(os.Args) > 1 {
		socket = os.Args[1]
	}

	c, err := qmp.Dial(socket)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("connect to %v ok\n", socket)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		var i int
		for {
			i++
			log.Printf("%v: start ping qga\n", i)
			c.SetDeadline(time.Now().Add(time.Duration(10*time.Second)))
			if err := c.Handshake(); err != nil {
				log.Printf("handshake error:%v\n", err)
				time.Sleep(time.Duration(time.Second))
				continue
			}
			break
		}
		wg.Done()
	}()
	wg.Wait()

	log.Printf("ping ok")

	defer c.Close()


}
//
//import (
//	"encoding/json"
//	"fmt"
//	"log"
//	"os"
//	"time"
//
//	"github.com/digitalocean/go-qemu/qmp"
//)
//
//
//type StatusResult struct {
//	ID     string `json:"id"`
//	Return struct {
//		Running    bool   `json:"running"`
//		Singlestep bool   `json:"singlestep"`
//		Status     string `json:"status"`
//	} `json:"return"`
//}
//
//func main() {
//	socket:="/run/vc/vm/19cc5ad191371184541dcc0c0ff197a3a59247cd76025a511f25aa6cc59bd897/qga.sock"
//	if len(os.Args) > 1 {
//		socket = os.Args[1]
//	}
//
//	fmt.Printf("start new socket monitor to %v\n", socket)
//	monitor, err := qmp.NewSocketMonitor("unix", socket, 2*time.Second)
//	if err != nil {
//		log.Fatalf("new socket monitor to %v error:%v", socket, err)
//	}
//
//	fmt.Printf("connect to monitor\n")
//	monitor.Connect()
//	defer monitor.Disconnect()
//
//
//	fmt.Printf("execute query-status\n")
//	cmd := []byte(`{ "execute": "query-status" }`)
//	raw, _ := monitor.Run(cmd)
//
//	fmt.Printf("process return of query-status")
//	var result StatusResult
//	json.Unmarshal(raw, &result)
//
//	fmt.Printf("result:%v" , result.Return.Status)
//}
//


