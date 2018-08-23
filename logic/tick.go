package logic

import (
	"time"

	"github.com/fyreek/Schmonk/actions"
	"github.com/fyreek/Schmonk/config"
	"github.com/fyreek/Schmonk/global"
)

func TickLoop() {
	tickTime := (time.Millisecond * 1000) / time.Duration(config.Config.Server.TickRate)
	dur := time.Duration(tickTime)
	ticker := time.NewTicker(dur)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				if global.GameActive {
					global.Mutex.Lock()
					l := len(global.Players)
					global.Mutex.Unlock()
					if l > 0 {
						actions.SendPlayerList(1)
					}
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
