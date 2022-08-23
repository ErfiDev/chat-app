package models

import "github.com/ErfiDev/chat-app/contract"

type Event struct {
	Type  int
	Uname string
	Rname string
	Conn  contract.WsConn
}
