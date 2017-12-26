package data

import (
  "fmt"
  ws "github.com/gorilla/websocket"
  gdax "github.com/preichenberger/go-gdax"
)

var address = "wss://ws-feed.gdax.com"

type GdaxReader struct {
  config gdax.Message
}

func NewReader(channel string, products []string) *GdaxReader {
  return &GdaxReader{
    config: gdax.Message{
      Type: "subscribe",
      Channels: []gdax.MessageChannel{
        gdax.MessageChannel{
          Name: channel,
          ProductIds: products,
        },
      },
    },
  }
}

func (w *GdaxReader) Listen() {
  var wsDialer ws.Dialer
  conn, _, err := wsDialer.Dial(address, nil)
  if err != nil {
    fmt.Println(err.Error())
  }
  if err := conn.WriteJSON(w.config); err != nil {
    panic(err)
  }

  var msg gdax.Message
  for true {
    if err := conn.ReadJSON(&msg); err != nil {
      // TODO(erich): replace with log
      panic(err)
    }
    // TODO(erich): probably want to report this on a channel
    fmt.Printf("%+v\n", msg)
  }
}
