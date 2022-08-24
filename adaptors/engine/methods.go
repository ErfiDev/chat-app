package engine

import (
	"fmt"
	"github.com/ErfiDev/chat-app/constant"
	"github.com/ErfiDev/chat-app/models"
)

func (en *Engine) handleEvent(e *models.Event) {
	switch e.Type {
	case constant.JoinEvent:
		room := en.findRoom(e.Rname)
		if room == nil {
			r := en.createRoom(e.Rname)
			usr := &models.User{
				Uname: e.Uname,
				Conn:  e.Conn,
			}

			r.Users = append(r.Users, usr)
		} else {
			usr := &models.User{
				Uname: e.Uname,
				Conn:  e.Conn,
			}
			room.Users = append(room.Users, usr)

			en.sendNotification(&models.SysMessage{
				Data:  fmt.Sprintf("user %s joined to room!", usr.Uname),
				Room:  room.Name,
				Uname: usr.Uname,
				Type:  constant.JoinEvent,
			})
		}

	case constant.LeaveEvent:
		room := en.findRoom(e.Rname)
		var users []*models.User

		for _, u := range room.Users {
			if u.Uname != e.Uname {
				users = append(users, u)
			}
		}

		room.Users = users
		en.sendNotification(&models.SysMessage{
			Data:  fmt.Sprintf("user %s leaved the room", e.Uname),
			Room:  room.Name,
			Uname: e.Uname,
			Type:  constant.LeaveEvent,
		})

	default:

	}
}

func (en *Engine) SendEvent(e *models.Event) {
	en.events <- e
}

func (en *Engine) SendMessage(m *models.Message) {
	en.messages <- m
}

func (en *Engine) broadcast(m *models.Message) {
	room := en.findRoom(m.Room)
	if room == nil {
		en.logger.Println("room not found: ", room.Name)
		return
	}

	for _, u := range room.Users {
		err := u.Conn.WriteJSON(m)
		if err != nil {
			en.logger.Printf("broadcast err: %s\n", err)
		}
	}
}

func (en *Engine) sendNotification(m *models.SysMessage) {
	room := en.findRoom(m.Room)
	if room == nil {
		en.logger.Println("room not found: ", room.Name)
		return
	}

	for _, u := range room.Users {
		err := u.Conn.WriteJSON(m)
		if err != nil {
			en.logger.Printf("broadcast sys message err: %s\n", err)
		}
	}
}

func (en *Engine) findRoom(name string) *models.Room {
	var room *models.Room
	for i, r := range en.rooms {
		if r.Name == name {
			room = en.rooms[i]
		}
	}

	return room
}

func (en *Engine) createRoom(name string) *models.Room {
	r := models.Room{
		Name:  name,
		Users: []*models.User{},
	}

	en.rooms = append(en.rooms, &r)

	return &r
}

func (en *Engine) Quit() {
	ch := make(chan error)
	en.quit <- ch
}
