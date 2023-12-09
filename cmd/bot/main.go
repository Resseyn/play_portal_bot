package main

import (
	"fmt"
	"net/http"
	"play_portal_bot/internal/botBase"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/databaseModels"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/internal/server"
	"play_portal_bot/pkg/database"
	"play_portal_bot/pkg/utils/structures"
	"time"
)

func main() {
	for k, v := range structures.Codes {
		fmt.Println(k+":", v)
	}
	fmt.Println("hello world")
	fmt.Println(time.Now().Format("02.01.2006 15:04"))
	loggers.InitLogger()
	database.InitDatabase()
	databaseModels.InitModels()
	for i := 0; i < 8; i++ {
		go func() {
			database.GlobalDatabase.Exec("INSERT INTO keys (key, key_code, avaliable) VALUES ($1, $2, $3)", helpingMethods.RandStringRunes(8), "app5", true)
		}()
	}
	mux := server.CreateMux()
	go http.ListenAndServe("localhost:8080", mux)
	err := botBase.BotStart()
	if err != nil {
		panic(err)
	}

}
