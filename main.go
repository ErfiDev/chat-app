package main

import (
	"github.com/ErfiDev/chat-app/adaptors/delivery"
	"github.com/ErfiDev/chat-app/adaptors/engine"
)

func main() {
	chatEngine := engine.New("./app.log")

	go chatEngine.Listener()

	delivery.Setup(chatEngine)
}
