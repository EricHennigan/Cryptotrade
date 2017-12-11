package data

import (
  "fmt"
  ws "github.com/gorilla/websocket"
  gdax "github.com/preichenberger/go-gdax"
)

type GdaxReader struct {
  conn *ws.Conn
  config map[string]string
}

func (w *GdaxReader) Init(address, product string) {
  var wsDialer ws.Dialer
  var err error
  w.conn, _, err = wsDialer.Dial(address, nil)
  if err != nil {
    fmt.Println(err.Error())
  }

  w.config = map[string]string{
    "type":       "subscribe",
    "product_id": product,
  }
}

func (w *GdaxReader) Listen() {
  if err := w.conn.WriteJSON(w.config); err != nil {
    panic(err)
  }

  var msg gdax.Message
  for true {
    if err := w.conn.ReadJSON(&msg); err != nil {
      // TODO(erich): replace with log
      panic(err)
    }
    // TODO(erich): probably want to report this on a channel
    fmt.Printf("%+v\n", msg)
  }
}
