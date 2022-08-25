package gui

import (
	"github.com/fasthttp/websocket"
	"github.com/jroimartin/gocui"
)

func New(uname, rname string, conn *websocket.Conn) (*Client, error) {
	gui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return &Client{}, err
	}

	cl := Client{
		Gui:   gui,
		uname: uname,
		rname: rname,
		conn:  conn,
	}

	return &cl, nil
}
