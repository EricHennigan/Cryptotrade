package main

import (
   "fmt"
   "github.com/Cryptotrade"
)

func main() {
  var w1 Worker{}
  w1.Init("wss://ws-feed.gdax.com", "subscribe", "BTC-USD")
}
