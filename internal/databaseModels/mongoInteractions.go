package databaseModels

//import (
//	"context"
//	"go.mongodb.org/mongo-driver/bson"
//	"play_portal_bot/pkg/database"
//	"play_portal_bot/pkg/utils/structures"
//)
//
//func GetAllOfMapsFromMongo() error {
//	pages := database.MongoDB.Database("play_bot_DB").Collection("pages")
//	handlers := database.MongoDB.Database("play_bot_DB").Collection("handlers")
//	parameters := database.MongoDB.Database("play_bot_DB").Collection("parameters")
//	commands := database.MongoDB.Database("play_bot_DB").Collection("commands")
//	prices := database.MongoDB.Database("play_bot_DB").Collection("prices")
//	codes := database.MongoDB.Database("play_bot_DB").Collection("codes")
//	cur, _ := pages.Find(context.Background(), bson.D{})
//	err := cur.Decode(&structures.Pages)
//	if err != nil {
//		return err
//	} //TODO: ну вот короче так все из монго достаешь и запихиваешь по словарям
//}
//
//func AddNewPageToMongo
