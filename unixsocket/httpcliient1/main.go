package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	containerId := flag.String("container-id", "04d33f3f6b019c081e3cecd217815d669136c979393ea02506091fcf83dceef4", "container id")
	flag.Parse()

	req, _ := http.NewRequest("GET", fmt.Sprintf("http://localhost/containers/%v/json", *containerId), nil)
	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout("unix", "/var/run/pouchd.sock", time.Second*5)
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * 5))
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * 5,
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("response:\n%v", string(buf))
}
