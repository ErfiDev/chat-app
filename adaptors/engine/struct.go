package engine

import "github.com/ErfiDev/chat-app/models"

type Engine struct {
	messages chan *models.Message
	events   chan *models.Event
	quit     chan chan error
	rooms    []*models.Room
}
