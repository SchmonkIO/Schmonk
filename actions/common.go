package actions

import (
	"github.com/fyreek/Schmonk/global"
	"github.com/fyreek/Schmonk/models"
)

func SendToAll(mt int, message []byte, player *models.Player) {
	global.Mutex.Lock()
	defer global.Mutex.Unlock()
	if len(global.Players) > 0 {
		for _, co := range global.Players {
			if co.ID != player.ID {
				co.Connection.WriteMessage(mt, message)
			}
		}
	}
}
