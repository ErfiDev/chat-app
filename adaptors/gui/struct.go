package gui

import (
	"github.com/fasthttp/websocket"
	"github.com/jroimartin/gocui"
)

type Client struct {
	*gocui.Gui
	uname string
	rname string
	conn  *websocket.Conn
}
