package helpers

import "github.com/gorilla/websocket"

func SendErrorMessage(ws *websocket.Conn, messageType int,err error) error {
	writeErr := ws.WriteMessage(messageType, []byte(err.Error()))
	if writeErr != nil {
		return writeErr
	}
	return nil
}
