package main

import (
	"fmt"
	"play_portal_bot/internal/botBase"
	"play_portal_bot/internal/loggers"
)

func main() {
	fmt.Println("hello world")
	loggers.InitLogger()
	err := botBase.BotStart()
	if err != nil {
		panic(err)
	}

}
