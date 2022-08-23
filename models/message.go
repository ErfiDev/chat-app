package models

type Message struct {
	From string `json:"from"`
	Room string `json:"room"`
	Data string `json:"data"`
}

type SysMessage struct {
	Data string `json:"data"`
}
