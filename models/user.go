package models

import (
	"github.com/gofiber/websocket/v2"
)

type User struct {
	ID    string
	Uname string
	Conn  *websocket.Conn
}
