package main

import (
	"fmt"
	"play_portal_bot/internal/loggers"
)

func main() {
	fmt.Println("hello world")
	loggers.InitLogger()
	loggers.ErrorLogger.Println("общий лог ура")
}
