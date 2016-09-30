package main

import (
  "golang.org/x/net/websocket"
  //"io"
  "net/http"
  "fmt"
)

func echoHandler(ws *websocket.Conn) {

  //UPDATE
  /*
  While the above example works, it’s probably better to keep the socket alive.
  In the above example the Golang server will keep the socket alive for one read and one write, and then exit.
  The following example keeps the handler alive in a loop:
  */
  for {
    receivedtext := make([]byte, 100)

    n,err := ws.Read(receivedtext)

    if err != nil {
      fmt.Printf("Received: %d bytes\n",n)
    }

    s := string(receivedtext[:n])
    fmt.Printf("Received: %d bytes: %s\n",n,s)

    // io.Copy(ws, ws)
    // fmt.Printf("Sent: %s\n",s)
  }
}

func main() {
  http.Handle("/echo", websocket.Handler(echoHandler))
  http.Handle("/", http.FileServer(http.Dir(".")))
  err := http.ListenAndServe(":8888", nil)
  if err != nil {
    panic("Error: " + err.Error())
  }
}
