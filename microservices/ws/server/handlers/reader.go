package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"ws/models"
	"ws/server/helpers"
	"ws/types"
)

type Reader struct {
	ctx  context.Context
	conn *websocket.Conn
	types.UserCrud
}

func NewReader(ctx context.Context, conn *websocket.Conn, crud types.UserCrud) *Reader {
	return &Reader{
		ctx:    ctx,
		conn:   conn,
		UserCrud: crud,
	}
}

func (r *Reader) Read() error {
	var req models.UserRequest
	//var userData ws.UserData
	for {
		messageType, data, err := r.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return err
		}
		err = json.Unmarshal(data, &req)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println(messageType)
		switch req.OpCode {
		case 1:
			fmt.Printf("\n\ncase 1:\nreq.Data[0]:\t%v\n\n", req.Data[0])
			err = r.Create(req.Data[0])
			if err != nil {
				err := helpers.SendErrorMessage(r.conn, messageType, err)
				if err != nil {
					log.Fatalf("%v error in sending response", err) // TODO error list
				}
			}
		case 2:
			user, err := r.Get(req.Data[0].ID)
			if err != nil {
				err := helpers.SendErrorMessage(r.conn, messageType, err)
				if err != nil {
					log.Fatalf("%v error in sending response", err) // TODO error list
				}
			}
			toJson, err := json.Marshal(user)
			if err != nil {
				err := helpers.SendErrorMessage(r.conn, messageType, err)
				if err != nil {
					log.Fatalf("%v error in marshaling response", err) // TODO error list
				}
			}
			err = r.conn.WriteMessage(messageType, toJson)
			if err != nil {
				log.Println(err)
			}
		case 3:
			err = r.Update(req.Data[0].ID, req.Data[0])
			if err != nil {
				err := helpers.SendErrorMessage(r.conn, messageType, err)
				if err != nil {
					log.Fatalf("%v error in sending response", err) // TODO error list
				}
			}
		case 4:
			err= r.Delete(req.Data[0].ID)
			if err != nil {
				err := helpers.SendErrorMessage(r.conn, messageType, err)
				if err != nil {
					log.Fatalf("%v error in sending response", err) // TODO error list
				}
			}
		}
	}
}
