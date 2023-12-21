package main

import (
	"fmt"
	"net/http"
	"play_portal_bot/internal/botBase"
	"play_portal_bot/internal/databaseModels"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/internal/server"
	"play_portal_bot/pkg/database"
	"play_portal_bot/pkg/utils/structures"
	"time"
)

func main() {
	//Pages = unmarshal from MONGO
	fuc := time.Now()
	//for i := 0; i < 10000000; i++ {
	//	structures.Commands["Success"] = strconv.FormatInt(int64(i), 10)
	//}//TODO: тест с редисом
	fmt.Println(time.Now().Sub(fuc))

	for k, v := range structures.Codes {
		fmt.Println(k+":", v)
	}
	fmt.Println("hello world")
	fmt.Println(time.Now().Format("02.01.2006 15:04"))
	loggers.InitLogger()
	database.InitDatabase()
	err := database.InitMongo()
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return
	}
	databaseModels.InitModels()

	//err = databaseModels.AddNewPageToMongo(&structures.TypicalPage{}, "", []string{}, []float64{}, map[string]string{}, []string{})
	//if err != nil {
	//	loggers.ErrorLogger.Println(err)
	//	return
	//}
	//helpingMethods.ParseMaps() //TODO: only for debugging database
	err = databaseModels.GetAllOfMapsFromMongo()
	if err != nil {
		loggers.ErrorLogger.Fatal(err)
		return
	}

	//for i := 0; i < 8; i++ {
	//	go func() {
	//		database.GlobalDatabase.Exec("INSERT INTO keys (key, key_code, avaliable) VALUES ($1, $2, $3)", helpingMethods.RandStringRunes(8), "app5", true)
	//	}()
	//}
	mux := server.CreateMux()
	go http.ListenAndServe("localhost:8080", mux)
	err = botBase.BotStart()
	if err != nil {
		panic(err)
	}

}
