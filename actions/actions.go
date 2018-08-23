package actions

import (
	"github.com/fyreek/Schmonk/models"
	"github.com/fyreek/Schmonk/util"
)

func ChooseAction(message []byte, mt int, player *models.Player) {
	if util.CheckAction(message, "startGame") {
		StartGame()
	}
	if util.CheckAction(message, "stopGame") {
		StopGame()
	}
	if util.CheckAction(message, "chat") {
		Chat(mt, message, player)
	}
	if util.CheckAction(message, "move") {
		Move(message, player)
	}
}
