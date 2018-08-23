package logic

import (
	"strconv"

	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"

	"github.com/fyreek/Schmonk/actions"
	"github.com/fyreek/Schmonk/global"
	"github.com/fyreek/Schmonk/models"
	"github.com/fyreek/Schmonk/util"
)

func PlayerSetup(c *websocket.Conn) {
	global.Mutex.Lock()
	util.LogToConsole("Connected Users: " + strconv.Itoa(len(global.Players)))
	global.Mutex.Unlock()
	player := models.Player{}
	player.ID = bson.NewObjectId()
	player.Connection = c
	player.Status = "unset"

	defer c.Close()
	playerLoop(c, &player)
}

func playerLoop(c *websocket.Conn, player *models.Player) {
	for {
		mt, message, err := c.ReadMessage()
		util.LogToConsole("MT: " + strconv.Itoa(mt))
		util.LogToConsole("Message: " + string(message))
		if err != nil {
			if mt == -1 {
				actions.Disconnect(mt, player)
				util.LogToConsole("disconnect player:", err)
			} else {
				util.LogToConsole("message read error:", err)
			}
			break
		}
		if player.Status == "unset" {
			if util.CheckAction(message, "join") {
				actions.Join(mt, message, player)
			} else {
				player.Connection.WriteMessage(mt, []byte("set name first"))
			}
		} else {
			actions.ChooseAction(message, mt, player)
		}
	}
}
