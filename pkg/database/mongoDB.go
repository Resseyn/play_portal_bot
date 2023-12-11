package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"play_portal_bot/internal/loggers"
	"time"
)

var MongoDB *mongo.Client
var MongoContext context.Context

func InitMongo() error {
	var err error
	MongoContext, _ = context.WithTimeout(context.Background(), 10*time.Second)
	MongoDB, err = mongo.Connect(MongoContext, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return err
	}
	err = MongoDB.Ping(MongoContext, readpref.Primary())
	return err
}
