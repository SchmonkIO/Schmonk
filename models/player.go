package models

import (
	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
)

type Player struct {
	ID         bson.ObjectId   `json:"_id"`
	Name       string          `json:"name"`
	PosX       float32         `json:"posx"`
	PosY       float32         `json:"posy"`
	Status     string          `json:"status"`
	Color      string          `json:"color"`
	Connection *websocket.Conn `json:"-"`
}
