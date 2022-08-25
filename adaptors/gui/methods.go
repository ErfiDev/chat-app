package gui

import (
	"encoding/json"
	"fmt"
	"github.com/ErfiDev/chat-app/constant"
	"github.com/ErfiDev/chat-app/dto"
	"github.com/jroimartin/gocui"
	"strings"
	"time"
)

func (c *Client) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	g.Cursor = true

	if messages, err := g.SetView("messages", 0, 0, maxX-20, maxY-5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		messages.Title = "messages"
		messages.Autoscroll = true
		messages.Wrap = true
	}

	if input, err := g.SetView("input", 0, maxY-5, maxX-20, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		input.Title = "send"
		input.Autoscroll = false
		input.Wrap = true
		input.Editable = true
	}

	if users, err := g.SetView("users", maxX-20, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		users.Title = "users"
		users.Autoscroll = false
		users.Wrap = true
	}
	g.SetCurrentView("input")

	return nil
}

func (c *Client) Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func (c *Client) SendMessage(g *gocui.Gui, v *gocui.View) error {
	if len(v.Buffer()) == 0 {
		v.SetCursor(0, 0)
		v.Clear()
		return nil
	}
	message := dto.Message{
		From: c.uname,
		Room: c.rname,
		Data: v.Buffer(),
	}

	err := c.conn.WriteJSON(&message)
	if err != nil {
		return err
	}

	v.SetCursor(0, 0)
	v.Clear()

	return nil
}

func (c *Client) ReceiveMsg() {
	for {
		msg := &dto.Message{}
		sysMsg := &dto.SysMessage{}

		_, bytes, err := c.conn.ReadMessage()
		if err != nil {
			return
		}

		err = json.Unmarshal(bytes, msg)
		if err != nil {
			return
		}
		err = json.Unmarshal(bytes, sysMsg)
		if err != nil {
			return
		}

		c.Update(func(g *gocui.Gui) error {
			if msg.From != "" {
				view, _ := c.View("messages")
				fmt.Fprint(view, fmt.Sprintf("%s: %s", time.Kitchen, msg.Data))

				return nil
			} else if sysMsg.Uname != "" {
				switch sysMsg.Type {
				case constant.JoinEvent:
					view, _ := c.View("messages")
					viewUsers, _ := c.View("users")
					fmt.Fprintf(view, sysMsg.Data)
					fmt.Fprintf(viewUsers, fmt.Sprintf("%s\n", sysMsg.Uname))
					return nil

				case constant.LeaveEvent:
					view, _ := c.View("messages")
					viewUsers, _ := c.View("users")
					fmt.Fprintf(view, sysMsg.Data)

					buf := viewUsers.Buffer()
					newBuf := strings.Replace(buf, sysMsg.Uname, "", 1)
					fmt.Fprintf(viewUsers, newBuf)
					return nil
				}

				return nil
			}

			return nil
		})
	}
}
