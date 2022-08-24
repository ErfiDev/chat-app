package models

import (
	"github.com/gofiber/websocket/v2"
)

type User struct {
	Uname string
	Conn  *websocket.Conn
}
