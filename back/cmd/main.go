package main

import (
	"fmt"

	ws "github.com/BigWaffleMonster/chat/pkg"
)

func main() {
	_ = ws.StartServer(messageHandler)
}

func messageHandler(msg ws.Message) {
	fmt.Print(msg)
}
