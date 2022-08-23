package engine

import (
	"github.com/ErfiDev/chat-app/models"
	"log"
)

func (en *Engine) handleEvent(e *models.Event) {

}

func (en *Engine) SendEvent(e *models.Event) {
	en.events <- e
}

func (en *Engine) SendMessage(m *models.Message) {
	en.messages <- m
}

func (en *Engine) broadcast(m *models.Message) {
	var room *models.Room
	for i, r := range en.rooms {
		if r.Name == m.Room {
			room = en.rooms[i]
		}
	}

	for _, u := range room.Users {
		err := u.Conn.WriteJSON(&m)
		log.Printf("We got err %s", err)
	}
}

func (en *Engine) sendNotification(m *models.SysMessage) {

}
