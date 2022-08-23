package tools

import (
	"io"
	"log"
)

func NewLogger(output io.Writer) *log.Logger {
	return log.New(output, "chat_app: ", log.Ldate|log.Llongfile)
}
