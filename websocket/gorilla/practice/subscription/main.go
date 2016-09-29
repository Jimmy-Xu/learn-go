package main

import (
    "github.com/gorilla/websocket"
    "net/http"
    "os"
    "fmt"
    "io/ioutil"
    "time"
    "encoding/json"
)

//First we’ll need to define our Person type:
type Person struct {
    Name string
    Age  int
}

//We’ll also need to create an upgraded variable, in which we define our read and write buffer sizes.
var upgrader = websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
}

func main() {
    indexFile, err := os.Open("html/index.html")
    if err != nil {
        fmt.Println(err)
    }
    index, err := ioutil.ReadAll(indexFile)
    if err != nil {
        fmt.Println(err)
    }
    http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
      //Now, how do we create the websocket connection? Pretty easily in fact:
      conn, err := upgrader.Upgrade(w, r, nil)
      if err != nil {
        fmt.Println(err)
        return
      }
      fmt.Println("Client subscribed")

      //Now let’s create Bill, our person, right after we get the client subscribed:
      myPerson := Person{
        Name: "Bill",
        Age:  0,
      }

      //Now we need the main websocket handling code, which we will wrap into an endless for loop,
      //which we get out of only if the channel closes or Bill gets 40 seconds old.
      for {
        time.Sleep(2 * time.Second)
        if myPerson.Age < 40 {
          myJson, err := json.Marshal(myPerson)
          if err != nil {
            fmt.Println(err)
            return
          }
          err = conn.WriteMessage(websocket.TextMessage, myJson)
          if err != nil {
            fmt.Println(err)
            break
          }
          myPerson.Age += 2
        } else {
          conn.Close()
          break
        }
      }
      fmt.Println("Client unsubscribed")

    })
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, string(index))
    })
    http.ListenAndServe(":3000", nil)
}
