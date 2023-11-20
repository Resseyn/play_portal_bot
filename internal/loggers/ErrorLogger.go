package loggers

import (
	"log"
	"os"
)

var ErrorLogger *log.Logger

func InitLogger() {
	file, err := os.OpenFile("/Users/romanovmaksim/GolandProjects/SnippetsTESTBYGUIDE/logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	ErrorLogger = log.New(file, "", log.LstdFlags)
}
