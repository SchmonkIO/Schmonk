package actions

import (
	"encoding/json"

	"github.com/fyreek/Schmonk/models"
	"github.com/fyreek/Schmonk/util"
)

func Chat(mt int, message []byte, player *models.Player) {
	data := models.ChatRequest{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		util.LogToConsole(err.Error())
		player.Connection.WriteMessage(mt, []byte("invalid json"))
	} else {
		resp := models.ChatRespose{}
		resp.Action = "chat"
		resp.Sender = player.ID
		resp.Message = data.Message
		chatJSON, _ := json.Marshal(resp)
		SendToAll(mt, chatJSON, player)
	}
}
