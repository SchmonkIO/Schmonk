package actions

import (
	"github.com/fyreek/Schmonk/global"
	"github.com/fyreek/Schmonk/models"
	"github.com/fyreek/Schmonk/util"
)

func Disconnect(mt int, player *models.Player) {
	global.Mutex.Lock()
	delete(global.Players, player.ID.Hex())
	global.Mutex.Unlock()
	SendPlayerList(mt)
	global.Mutex.Lock()
	l := len(global.Players)
	global.Mutex.Unlock()
	util.LogToConsole("Connected Players:", l)
}
