package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Room struct {
	ID       bson.ObjectId `json:"_id"`
	Name     string        `json:"name"`
	Password string        `json:"password"`
	Slots    int           `json:"slots"`
	Owner    bson.ObjectId `json:"owner"`
}

// func (r *Room) create() error {
// 	global.Rooms[r.ID] = r
// 	return nil
// }
