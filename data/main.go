package main

import (
	"fmt"
	ws "github.com/gorilla/websocket"
	gdax "github.com/preichenberger/go-gdax"
)

func main() {
	var wsDialer ws.Dialer
	wsConn, _, err := wsDialer.Dial("wss://ws-feed.gdax.com", nil)
	if err != nil {
		println(err.Error())
	}

	subscribe := map[string]string{
		"type":       "subscribe",
		"product_id": "BTC-USD",
	}
	if err := wsConn.WriteJSON(subscribe); err != nil {
		println(err.Error())
	}

	count := 0
	message := gdax.Message{}
	for true {
		if err := wsConn.ReadJSON(&message); err != nil {
			println(err.Error())
			break
		}
		fmt.Printf("%d %+v\n", count, message)
		count++

		if message.Type == "match" {
			println("Got a match")
		}
	}
}
