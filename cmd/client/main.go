package main

import (
	"flag"
	"fmt"
	"github.com/ErfiDev/chat-app/adaptors/gui"
	"github.com/ErfiDev/chat-app/constant"
	"github.com/fasthttp/websocket"
	"github.com/google/uuid"
	"github.com/jroimartin/gocui"
	"log"
)

var uu = uuid.New()

func main() {
	var uname, rname string
	flag.StringVar(&uname, "username", "", "username for connecting")
	flag.StringVar(&rname, "room", "", "room name for connecting")

	flag.Parse()

	if uname == "" || rname == "" {
		log.Fatalln("username and room flag required")
	}

	id := uu.String()

	url := fmt.Sprintf("ws://localhost%s/connect/%s/%s/%s", constant.PORT, rname, uname, id)

	conn, resp, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatalf("handshake error: %s", err)
	}

	defer conn.Close()
	defer resp.Body.Close()

	ui, err := gui.New(uname, rname, id, conn)
	if err != nil {
		log.Fatalf("gui init error: %s", err)
	}

	ui.SetManagerFunc(ui.Layout)
	if err := ui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, ui.Quit); err != nil {
		log.Fatalln(err)
	}
	if err := ui.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, ui.SendMessage); err != nil {
		log.Fatalln(err)
	}

	go ui.ReceiveMsg()

	if err = ui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatalln(err)
	}
}
