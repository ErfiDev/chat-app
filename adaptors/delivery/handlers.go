package delivery

import (
	"encoding/json"
	"github.com/ErfiDev/chat-app/constant"
	"github.com/ErfiDev/chat-app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Connect() fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		uname := c.Params("uname")
		room := c.Params("room")

		joinReq := models.Event{
			Type:  constant.JoinEvent,
			Uname: uname,
			Rname: room,
			Conn:  c,
		}

		chatEngine.SendEvent(&joinReq)

		for {
			t, bytes, err := c.ReadMessage()
			if err != nil {
				return
			}

			switch t {
			case websocket.CloseMessage:
				event := models.Event{
					Type:  constant.LeaveEvent,
					Uname: uname,
					Rname: room,
					Conn:  c,
				}

				chatEngine.SendEvent(&event)
				return

			default:
				var msg models.Message
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
