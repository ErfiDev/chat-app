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

		joinReq := &dto.Event{
			Type:  constant.JoinEvent,
			Rname: room,
			User: &models.User{
				Uname: uname,
				Conn:  c,
			},
		}

		chatEngine.SendEvent(joinReq)

		for {
			t, bytes, err := c.ReadMessage()
			if err != nil {
				return
			}

			switch t {
			case websocket.CloseMessage:
				event := &dto.Event{
					Type:  constant.LeaveEvent,
					Rname: room,
					User: &models.User{
						Uname: uname,
						Conn:  c,
					},
				}

				chatEngine.SendEvent(event)
				return

			default:
				var msg dto.Message
				err := json.Unmarshal(bytes, &msg)
				if err != nil {
					return
				}

				if msg.From != "" {
					if msg.From == uname {
						chatEngine.SendMessage(&msg)
					}
				}
			}
		}
	})
}
