package engine

import (
	"github.com/ErfiDev/chat-app/dto"
	"github.com/ErfiDev/chat-app/models"
	"github.com/ErfiDev/chat-app/tools"
)

func New(logOutput string) *Engine {
	file := tools.LoadFile(logOutput)

	e := &Engine{
		messages: make(chan *dto.Message),
		events:   make(chan *dto.Event),
		quit:     make(chan chan error),
		rooms:    make([]*models.Room, 0),
		logger:   tools.NewLogger(file),
	}

	return e
}
