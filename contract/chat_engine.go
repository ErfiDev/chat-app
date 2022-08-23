package contract

import "github.com/ErfiDev/chat-app/models"

type ChatEngine interface {
	handleEvent(e *models.Event)
	SendEvent(e *models.Event)
	SendMessage(m *models.Message)
	Listener()
	broadcast(m *models.Message)
	sendNotification(m *models.SysMessage)
}
