package engine

import (
	"github.com/ErfiDev/chat-app/dto"
	"github.com/ErfiDev/chat-app/models"
	"log"
)

type Engine struct {
	messages chan *dto.Message
	events   chan *dto.Event
	quit     chan chan error
	rooms    []*models.Room
	logger   *log.Logger
}
