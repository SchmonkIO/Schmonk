package util

import (
	"fmt"

	"github.com/fyreek/Schmonk/config"
)

func LogToConsole(a ...interface{}) {
	if config.Config.Server.Debug {
		fmt.Println(a)
	}
}
