package main

import (
	"fmt"
	"net/http"
	"play_portal_bot/internal/botBase"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/internal/server"
)

func main() {
	fmt.Println("hello world")
	loggers.InitLogger()
	mux := server.CreateMux()
	go http.ListenAndServe("localhost:8080", mux)
	err := botBase.BotStart()
	if err != nil {
		panic(err)
	}

}
