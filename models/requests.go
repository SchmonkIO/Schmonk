package models

type NameRequest struct {
	Action string `json:"action"`
	Name   string `json:"name"`
}

type ChatRequest struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}

type MoveRequest struct {
	Action string  `json:"action"`
	PosX   float32 `json:"posx,string"`
	PosY   float32 `json:"posy,string"`
}
