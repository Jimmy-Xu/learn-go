package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
	"golang.org/x/crypto/ssh"
)

//usage:
// ./scp -p xxxxxx main.go xjimmy@172.16.87.134:/home/xjimmy/main.go
// ./scp -p xxxxxx main.go xjimmy@172.16.87.134:~/main1.go

func main() {
	password := flag.String("p", "", "password")
	flag.Parse()

	if len(flag.Args()) != 2 {
		log.Fatal("please specify args")
	}

	var (
		localFile     string
		username      string
		serverIP      string
		remotePath    string
		localToRemote bool
	)

	if strings.Contains(flag.Args()[0], "@") && !strings.Contains(flag.Args()[1], "@") {
		//localFile = flag.Args()[1]
		//username, serverIP, remotePath = parseRemoteArgs(flag.Args()[0])
		//localToRemote = false
		log.Fatalf("not support copy from remote to local")
	} else if !strings.Contains(flag.Args()[0], "@") && strings.Contains(flag.Args()[1], "@") {
		localFile = flag.Args()[0]
		username, serverIP, remotePath = parseRemoteArgs(flag.Args()[1])
		localToRemote = true
	} else {
		log.Fatalf("support copy local file tto remote only")
	}

	fmt.Printf("localFile:%s username:%s serverIP:%s remotePath:%s password:%s localToRemoet:%v\n",
		localFile, username, serverIP, remotePath, *password, localToRemote)

	// Use SSH key authentication from the auth package
	// we ignore the host key in this example, please change this if you use this library
	//clientConfig, _ := auth.PrivateKey("username", "/path/to/rsa/key", ssh.InsecureIgnoreHostKey())
	clientConfig, _ := auth.PasswordKey(username, *password, ssh.InsecureIgnoreHostKey())

	// For other authentication methods see ssh.ClientConfig and ssh.AuthMethod

	// Create a new SCP client
	client := scp.NewClient(serverIP+":22", &clientConfig)

	// Connect to the remote server
	err := client.Connect()
	if err != nil {
		fmt.Println("Couldn't establish a connection to the remote server ", err)
		return
	}

	// Open a file
	//f, _ := os.Open(localFile)
	f, err := os.OpenFile(localFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Failed to open local file %s, error:%v", localFile, err)
		return
	}

	// Close client connection after the file has been copied
	defer client.Close()

	// Close the file after it has been copied
	defer f.Close()

	// Finaly, copy the file over
	// Usage: CopyFile(fileReader, remotePath, permission)

	if localToRemote {
		fmt.Printf("copy local(%s) to remote(%s)\n", localFile, remotePath)
		err = client.CopyFile(f, remotePath, "0655")
	} else {
		fmt.Printf("copy remote(%s) to local(%s)\n", remotePath, localFile)
		err = client.CopyFromFile(*f, remotePath, "0655")
	}

	if err != nil {
		fmt.Println("Error while copying file ", err)
	}
}

func parseRemoteArgs(remoteArgs string) (string, string, string) {
	//parse username
	tmp := strings.Split(remoteArgs, "@")
	if len(tmp) != 2 {
		log.Fatal("missing @")
	}
	username := tmp[0]

	//parse serverIP and remote Path
	tmp = strings.Split(tmp[1], ":")
	if len(tmp) != 2 {
		log.Fatal("missing :")
	}
	serverIP := tmp[0]
	remotePath := tmp[1]

	return username, serverIP, remotePath
}
