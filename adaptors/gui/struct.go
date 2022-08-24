package gui

import (
	"github.com/ErfiDev/chat-app/contract"
	"github.com/jroimartin/gocui"
)

type Client struct {
	*gocui.Gui
	uname string
	rname string
	conn  contract.WsConn
}
