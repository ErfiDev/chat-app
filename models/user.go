package models

import "github.com/ErfiDev/chat-app/contract"

type User struct {
	Uname string
	Conn  contract.WsConn
}
