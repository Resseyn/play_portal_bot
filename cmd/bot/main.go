package main

import (
	"fmt"
	"net/http"
	"play_portal_bot/internal/botBase"
	"play_portal_bot/internal/databaseModels"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/internal/server"
	"play_portal_bot/pkg/database"
)

func main() {
	fmt.Println("hello world")
	loggers.InitLogger()
	database.InitDatabase()
	databaseModels.InitModels()
	mux := server.CreateMux()
	go http.ListenAndServe("localhost:8080", mux)
	err := botBase.BotStart()
	if err != nil {
		panic(err)
	}

}
