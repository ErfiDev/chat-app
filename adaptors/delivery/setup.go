package delivery

import (
	"flag"
	"github.com/ErfiDev/chat-app/adaptors/engine"
	"github.com/gofiber/fiber/v2"
)

var chatEngine *engine.Engine

func Setup(en *engine.Engine) {
	chatEngine = en

	var port string
	flag.StringVar(&port, "port", ":5000", "Application listening port")

	flag.Parse()

	r := fiber.New()

	r.Get("/connect/:room/:uname", Connect())

	err := r.Listen(port)
	if err != nil {
		panic(err)
	}
}
