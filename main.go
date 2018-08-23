package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/rs/cors"

	"github.com/fyreek/Schmonk/config"
	"github.com/fyreek/Schmonk/logic"
	"github.com/fyreek/Schmonk/web"
)

func main() {
	setup()
	sAddress := config.Config.Server.IP + ":" + strconv.Itoa(config.Config.Server.Port)
	logic.TickLoop()
	setup()
	log.Fatal(http.ListenAndServe(sAddress, setupHandler()))
}

func setupHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", logic.Register)
	mux.HandleFunc("/", web.Load)
	//mux.Handle("/game", http.FileServer(http.Dir("./web")))
	mux.Handle("/game/", http.StripPrefix("/game/", http.FileServer(http.Dir("./web"))))
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            config.Config.Server.Debug,
	}).Handler(mux)
	return handler
}

func setup() {
	err := config.ReadConfig("server.conf")
	if err != nil {
		fmt.Println("[Failure] Could not read config file")
		os.Exit(1)
	}
	log.SetFlags(0)
}
