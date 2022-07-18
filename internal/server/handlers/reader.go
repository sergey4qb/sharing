package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"tr_redis_ws/internal/models"
	"tr_redis_ws/internal/storage"
)

type Reader struct {
	conn *websocket.Conn
	*storage.Redis
}

func NewReader(conn *websocket.Conn, redis *storage.Redis) *Reader {
	//repo := storage.New()
	//newRedis := storage.NewRedis(repo.Client, time.Hour)
	return &Reader{
		conn:    conn,
		Redis: redis,
	}
}

func (r *Reader) Read() error {
	//ws_user := new(models.User)
	req := new(models.Request)
	for {
		messageType, data, err := r.conn.ReadMessage()
		if err != nil {
			return err
		}
		err = json.Unmarshal(data, &req)
		if err != nil {
			return err
		}
		switch req.OpCode {
		case 1:
			err := r.Create(req.Data[0])
			if err != nil {
				return err
			}
			if err := r.conn.WriteMessage(messageType, data); err != nil {
				return err
			}

		case 2:
			fmt.Println("case 2")
			user, err := r.GetUser(req.Data[0].ID)
			if err != nil {
				return err
			}
			fmt.Println(user)

		}

		fmt.Println(req.Data)
		if err := r.conn.WriteMessage(messageType, data); err != nil {
			return err
		}
	}
}
