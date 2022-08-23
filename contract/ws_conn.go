package contract

type WsConn interface {
	Close() error
	WriteMessage(messageType int, data []byte) error
	WriteJSON(v interface{}) error
	ReadJSON(v interface{}) error
	ReadMessage() (messageType int, p []byte, err error)
}
