package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/fyreek/Schmonk/config"
	"github.com/fyreek/Schmonk/logic"
)

func main() {
	setup()
	sAddress := config.Config.Server.IP + ":" + strconv.Itoa(config.Config.Server.Port)
	logic.TickLoop()
	router := gin.Default()
	router.Static("/js", "./web/js")
	router.LoadHTMLGlob("web/*.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	router.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.html", gin.H{
			"title": "Main website",
		})
	})
	router.GET("/ws", func(c *gin.Context) {
		logic.Register(c.Writer, c.Request)
	})
	log.Fatal(router.Run(sAddress))
}

func setup() {
	err := config.ReadConfig("server.conf")
	if err != nil {
		fmt.Println("[Failure] Could not read config file")
		os.Exit(1)
	}
	log.SetFlags(0)
}
