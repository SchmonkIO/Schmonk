package models

import (
	"gopkg.in/mgo.v2/bson"
)

type PlayerList struct {
	Action  string             `json:"action"`
	ID      bson.ObjectId      `json:"_id"`
	Players map[string]*Player `json:"players"`
}

type ChatRespose struct {
	Action  string        `json:"action"`
	Sender  bson.ObjectId `json:"id"`
	Message string        `json:"message"`
}
