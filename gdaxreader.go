package main

import (
	"./data"
)

func main() {
	w1 := data.NewReader("level2", []string{"BTC-USD"})
	w1.Listen()
}
