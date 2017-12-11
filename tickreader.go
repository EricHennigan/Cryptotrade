package ctrade

import (
  "fmt"
  "net"
  ws "github.com/gorilla/websocket"
  gdax "github.com/preichenberger/go-gdax"
)

type Worker struct {
  Name string
  Config map[string]string
  Connection *Conn
  MessageChannel map[string]string
  Message gdax.Message{}
  Count int
}

func (wkr *Worker) Init(address string, action string, product, string){
  var wsDialer ws.Dialer
  wkr.Connection, _, err := wsDialer.Dialer.Dial(address, nil)
  if err != nil {
    fmt.Println(err.Error())
  }

  wkr.MessageChannel := map[string]string{
    "type":       action,
    "product_id": product,
  }
}

func (wkr *Worker) DoWork() {
  if err := wkr.Connection.WriteJSON(wkr.MessageChannel); err != nil {
    panic(err)
  }

  for true {
    if err := wkr.Connection.ReadJSON(&wkr.Message); err != nil {
      panic(err)
    }
    fmt.Printf("%d %+v\n", count, message)
    wkr.Count++
  }  
}

// package main

// import (
//  "fmt"
//  ws "github.com/gorilla/websocket"
//  gdax "github.com/preichenberger/go-gdax"
// )

// func main() {
//  var wsDialer ws.Dialer
//  wsConn, _, err := wsDialer.Dial("wss://ws-feed.gdax.com", nil)
//  if err != nil {
//    println(err.Error())
//  }

//  subscribe := map[string]string{
//    "type":       "subscribe",
//    "product_id": "BTC-USD",
//  }
//  if err := wsConn.WriteJSON(subscribe); err != nil {
//    println(err.Error())
//  }

//  count := 0
//  message := gdax.Message{}
//  for true {
//    if err := wsConn.ReadJSON(&message); err != nil {
//      println(err.Error())
//      break
//    }
//    fmt.Printf("%d %+v\n", count, message)
//    count++

//    if message.Type == "match" {
//      println("Got a match")
//    }
//  }
// }
