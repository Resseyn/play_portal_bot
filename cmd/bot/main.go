package main

import (
	"fmt"
	"net/http"
	"play_portal_bot/internal/botBase"
	"play_portal_bot/internal/databaseModels"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/internal/server"
	"play_portal_bot/pkg/database"
	"time"
)

func main() {
	//TODO: func show codes - for key show value in codes
	fmt.Println("hello world")
	fmt.Println(time.Now().Format("02.01.2006 15:04"))
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
