package actions

import (
	"encoding/json"

	"github.com/fyreek/Schmonk/global"
	"github.com/fyreek/Schmonk/models"
	"github.com/fyreek/Schmonk/util"
)

func SendPlayerList(mt int) {
	global.Mutex.Lock()
	l := len(global.Players)
	global.Mutex.Unlock()
	if l > 0 {
		if mt == -1 {
			mt = 1
		}
		global.Mutex.Lock()
		defer global.Mutex.Unlock()
		for _, co := range global.Players {
			pList := models.PlayerList{}
			pList.Action = "tick"
			pList.ID = co.ID
			pList.Players = global.Players
			pListJSON, _ := json.Marshal(pList)
			err := co.Connection.WriteMessage(mt, pListJSON)
			if err != nil {
				util.LogToConsole(err.Error())
			}
		}
	}
}
