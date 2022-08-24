package dto

import "github.com/ErfiDev/chat-app/models"

type Event struct {
	Type  int
	Rname string
	User  *models.User
}
