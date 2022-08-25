package engine

import (
	"fmt"
	"github.com/ErfiDev/chat-app/constant"
	"github.com/ErfiDev/chat-app/dto"
	"github.com/ErfiDev/chat-app/models"
)

func (en *Engine) handleEvent(e *dto.Event) {
	switch e.Type {
	case constant.JoinEvent:
		room := en.findRoom(e.Rname)
		if room == nil {
			r := en.createRoom(e.Rname)
			r.Users = append(r.Users, e.User)

			en.sendNotification(&dto.SysMessage{
				Data:  fmt.Sprintf("user %s joined to room!\n", e.User.Uname),
				Room:  e.Rname,
				Uname: e.User.Uname,
				Type:  constant.JoinEvent,
			})
		} else {
			ok := en.findUserInRoom(e.Rname, e.User.Uname)
			if !ok {
				room.Users = append(room.Users, e.User)

				en.sendNotification(&dto.SysMessage{
					Data:  fmt.Sprintf("user %s joined to room!\n", e.User.Uname),
					Room:  e.Rname,
					Uname: e.User.Uname,
					Type:  constant.JoinEvent,
				})
			}
		}

	case constant.LeaveEvent:
		room := en.findRoom(e.Rname)
		var users []*models.User

		for _, u := range room.Users {
			if u.Uname != e.User.Uname {
				users = append(users, u)
			}
		}

		room.Users = users
		en.sendNotification(&dto.SysMessage{
			Data:  fmt.Sprintf("user %s leaved the room!\n", e.User.Uname),
			Room:  room.Name,
			Uname: e.User.Uname,
			Type:  constant.LeaveEvent,
		})

	}
}

func (en *Engine) SendEvent(e *dto.Event) {
	en.events <- e
}

func (en *Engine) SendMessage(m *dto.Message) {
	en.messages <- m
}

func (en *Engine) broadcast(m *dto.Message) {
	room := en.findRoom(m.Room)
	if room == nil {
		en.logger.Println("Room not found: ", m.Room)
		return
	}

	for _, u := range room.Users {
		err := u.Conn.WriteJSON(m)
		if err != nil {
			en.logger.Printf("broadcast err: %s\n", err)
		}
	}
}

func (en *Engine) sendNotification(m *dto.SysMessage) {
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

func (en *Engine) findUserInRoom(room string, uname string) bool {
	r := en.findRoom(room)
	if r == nil {
		return false
	}

	for _, u := range r.Users {
		if u.Uname == uname {
			return true
		}
	}

	return false
}
