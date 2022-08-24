package models

type Message struct {
	From string `json:"from"`
	Room string `json:"room"`
	Data string `json:"data"`
}

type SysMessage struct {
	Data  string `json:"data"`
	Room  string `json:"room"`
	Uname string `json:"uname"`
	Type  int    `json:"type"`
}
