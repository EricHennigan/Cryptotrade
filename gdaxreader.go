package main

import (
   "./data"
)

func main() {
  w1 := data.GdaxReader{}
  w1.Init("wss://ws-feed.gdax.com", "BTC-USD")
  w1.Listen()
}
