package actions

import (
	"encoding/json"

	"github.com/fyreek/Schmonk/global"
	"github.com/fyreek/Schmonk/models"
	"github.com/fyreek/Schmonk/util"
)

func Join(mt int, message []byte, player *models.Player) {
	data := models.NameRequest{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		util.LogToConsole(err.Error())
		player.Connection.WriteMessage(mt, []byte("invalid json"))
	} else {
		player.Name = data.Name
		player.Color = util.GetRandomColor()
		player.Status = "lobby"
		global.Mutex.Lock()
		global.Players[player.ID.Hex()] = player
		global.Mutex.Unlock()
		util.LogToConsole(player.ID.Hex())
		global.Mutex.Lock()
		util.LogToConsole(global.Players[player.ID.Hex()])
		global.Mutex.Unlock()
		SendPlayerList(mt)
		global.Mutex.Lock()
		l := len(global.Players)
		global.Mutex.Unlock()
		util.LogToConsole("Connected Players:", l)
	}
}
