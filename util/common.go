package util

import "github.com/fyreek/Schmonk/models"

func Find(a []*models.Player, x *models.Player) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}
