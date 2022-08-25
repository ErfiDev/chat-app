package delivery

import (
	"encoding/json"
	"github.com/ErfiDev/chat-app/adaptors/engine"
	"github.com/ErfiDev/chat-app/constant"
	"github.com/ErfiDev/chat-app/dto"
	"github.com/ErfiDev/chat-app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Connect(chatEngine *engine.Engine) fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		uname := c.Params("uname")
		room := c.Params("room")
		id := c.Params("id")

		joinReq := &dto.Event{
			Type:  constant.JoinEvent,
			Rname: room,
			User: &models.User{
				Uname: uname,
				Conn:  c,
				ID:    id,
			},
		}

		chatEngine.SendEvent(joinReq)

		for {
			_, bytes, err := c.ReadMessage()
			if err != nil {
				chatEngine.SendEvent(&dto.Event{
					Type:  constant.LeaveEvent,
					Rname: room,
					User: &models.User{
						ID:    id,
						Uname: uname,
						Conn:  c,
					},
				})
				return
			}

			var msg dto.Message
			_ = json.Unmarshal(bytes, &msg)

			if msg.From == uname {
				chatEngine.SendMessage(&msg)
			}
		}
	})
}
