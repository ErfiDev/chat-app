package main

import (
	"github.com/ErfiDev/chat-app/adaptors/delivery"
	"github.com/ErfiDev/chat-app/adaptors/engine"
)

func main() {
	chatEngine := engine.New("./engine.log")

	go chatEngine.Listener()

	delivery.Setup(chatEngine)
}
