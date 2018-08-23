package actions

import (
	"encoding/json"

	"github.com/fyreek/Schmonk/global"
	"github.com/fyreek/Schmonk/models"
)

func Move(message []byte, player *models.Player) {
	data := models.MoveRequest{}
	err := json.Unmarshal(message, &data)
	if err == nil {
		global.Mutex.Lock()
		p := global.Players[player.ID.Hex()]
		global.Mutex.Unlock()
		if p != nil {
			p.PosX = data.PosX
			p.PosY = data.PosY
			global.Mutex.Lock()
			global.Players[player.ID.Hex()] = p
			global.Mutex.Unlock()
		}
	} else {
		print(err.Error())
	}
}
