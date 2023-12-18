package loggers

import (
	"log"
	"os"
)

var ErrorLogger *log.Logger
var GlobalLogger *log.Logger

func InitLogger() {
	file, err := os.OpenFile("logs/errlog.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	ErrorLogger = log.New(file, "", log.LstdFlags)
	file, err = os.OpenFile("logs/globallog.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	GlobalLogger = log.New(file, "", log.LstdFlags)
}
