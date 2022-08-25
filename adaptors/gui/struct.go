package gui

import (
	"github.com/fasthttp/websocket"
	"github.com/jroimartin/gocui"
)

type Client struct {
	*gocui.Gui
	id    string
	uname string
	rname string
	conn  *websocket.Conn
}
