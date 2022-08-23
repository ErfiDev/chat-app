package models

type Message struct {
	From string `json:"from"`
	Data string `json:"data"`
}

type SysMessage struct {
	Data string `json:"data"`
}
