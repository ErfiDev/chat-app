package contract

import (
	"github.com/ErfiDev/chat-app/dto"
)

type ChatEngine interface {
	handleEvent(e *dto.Event)
	SendEvent(e *dto.Event)
	SendMessage(m *dto.Message)
	Listener()
	broadcast(m *dto.Message)
	sendNotification(m *dto.SysMessage)
	Quit()
}
