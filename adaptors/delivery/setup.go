package delivery

import (
	"github.com/ErfiDev/chat-app/adaptors/engine"
	"github.com/ErfiDev/chat-app/constant"
	"github.com/gofiber/fiber/v2"
)

var chatEngine *engine.Engine

func Setup(en *engine.Engine) {
	chatEngine = en

	r := fiber.New()

	r.Get("/connect/:room/:uname", Connect())

	err := r.Listen(constant.PORT)
	if err != nil {
		panic(err)
	}
}
