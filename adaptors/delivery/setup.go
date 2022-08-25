package delivery

import (
	"github.com/ErfiDev/chat-app/adaptors/engine"
	"github.com/ErfiDev/chat-app/constant"
	"github.com/gofiber/fiber/v2"
)

func Setup(en *engine.Engine) {
	r := fiber.New()

	r.Get("/connect/:room/:uname", Connect(en))

	err := r.Listen(constant.PORT)
	if err != nil {
		panic(err)
	}
}
