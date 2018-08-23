package global

import (
	"sync"

	"github.com/fyreek/Schmonk/models"
)

var Players map[string]*models.Player = map[string]*models.Player{}
var Rooms map[string]*models.Room = map[string]*models.Room{}
var GameActive = false
var Mutex = &sync.Mutex{}
