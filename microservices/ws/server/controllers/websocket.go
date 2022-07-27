package controllers

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"ws/server/handlers"
)

func (c *Controllers) WebSocket(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	reader := handlers.NewReader(c.ctx, conn, c.grpc)
	err = reader.Read()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
