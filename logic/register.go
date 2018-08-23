package logic

import (
	"net/http"

	"github.com/fyreek/Schmonk/config"
	"github.com/fyreek/Schmonk/global"
	"github.com/fyreek/Schmonk/util"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Register(w http.ResponseWriter, r *http.Request) {
	if !config.Config.Server.CORS {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	}
	c, err := upgrader.Upgrade(w, r, nil)
	util.LogToConsole("upgrade request")
	if err != nil {
		util.LogToConsole("upgrade:", err)
		return
	}
	global.Mutex.Lock()
	l := len(global.Players)
	global.Mutex.Unlock()
	if l < config.Config.Server.Slots {
		PlayerSetup(c)
	} else {
		c.WriteMessage(1, []byte("Slots exceeded"))
		c.Close()
		return
	}
}
