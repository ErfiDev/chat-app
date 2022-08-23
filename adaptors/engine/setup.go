package engine

import (
	"github.com/ErfiDev/chat-app/models"
	"github.com/ErfiDev/chat-app/tools"
)

func New(logOutput string) *Engine {
	file := tools.LoadFile("./app.log")

	e := Engine{
		messages: make(chan *models.Message),
		events:   make(chan *models.Event),
		quit:     make(chan chan error),
		rooms:    []*models.Room{},
		logger:   tools.NewLogger(file),
	}

	return &e
}
