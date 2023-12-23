package databaseModels

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/database"
	"play_portal_bot/pkg/utils/structures"
)

func GetAllOfMapsFromMongo() error {
	pages := database.MongoDB.Database("play_bot_DB").Collection("pages")
	handlers := database.MongoDB.Database("play_bot_DB").Collection("handlers")
	parameters := database.MongoDB.Database("play_bot_DB").Collection("parameters")
	commands := database.MongoDB.Database("play_bot_DB").Collection("commands")
	prices := database.MongoDB.Database("play_bot_DB").Collection("prices")
	codes := database.MongoDB.Database("play_bot_DB").Collection("codes")
	futurePages := make(map[string]*structures.TypicalPage)
	err := parseMongoMap(pages, futurePages)
	if err != nil {
		panic(err)
		return err
	}
	structures.Pages = futurePages

	futureHandlers := make(map[string]string)
	err = parseMongoMap(handlers, futureHandlers)
	if err != nil {
		panic(err)
		return err
	}
	structures.Handlers = futureHandlers

	futureParameters := make(map[string][]string)
	err = parseMongoMap(parameters, futureParameters)
	if err != nil {
		panic(err)
		return err
	}
	structures.Parameters = futureParameters

	futureCommands := make(map[string]string)
	err = parseMongoMap(commands, futureCommands)
	if err != nil {
		panic(err)
		return err
	}
	structures.Commands = futureCommands

	futurePrices := make(map[string]float64)
	err = parseMongoMap(prices, futurePrices)
	if err != nil {
		panic(err)
		return err
	}
	structures.Prices = futurePrices

	futureCodes := make(map[string]string)
	err = parseMongoMap(codes, futureCodes)
	if err != nil {
		panic(err)
		return err
	}
	structures.Codes = futureCodes

	return nil
}

func AddNewPageToMongo(page *structures.TypicalPage, goodPages []*structures.TypicalPage, mainCommandName string, handlerName string, handlerParams []string, prices []float64, codesTexts map[string]string) error {
	ctx := context.Background()

	// Добавление страницы
	pages := database.MongoDB.Database("play_bot_DB").Collection("pages")
	err := editPages(pages, ctx, page, goodPages, mainCommandName)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return err
	}
	// Добавление обработчика
	handlers := database.MongoDB.Database("play_bot_DB").Collection("handlers")
	filter := makeFilter(handlers)
	for _, good := range page.Goods {
		_, err = handlers.UpdateOne(ctx, filter, bson.M{"$set": map[string]string{good.Custom: handlerName}})
		if err != nil {
			return err
		}
	}

	// Добавление параметров
	if handlerName != "" {
		parameters := database.MongoDB.Database("play_bot_DB").Collection("parameters")
		filter = makeFilter(parameters)
		_, err = parameters.UpdateOne(ctx, filter, bson.M{"$set": map[string][]string{handlerName: handlerParams}})
		if err != nil {
			return err
		}
	}

	// Добавление команд
	commandsColl := database.MongoDB.Database("play_bot_DB").Collection("commands")
	filter = makeFilter(commandsColl)
	for _, commands := range page.Commands {
		for _, command := range commands {
			_, err = commandsColl.UpdateOne(ctx, filter, bson.M{"$set": map[string]string{command.Text: command.Command}})
			if err != nil {
				return err
			}
		}

	}
	_, err = commandsColl.UpdateOne(ctx, filter, bson.M{"$set": map[string]string{mainCommandName: page.MainCommand}})
	if err != nil {
		return err
	}
	// Добавление цен
	pricesCollection := database.MongoDB.Database("play_bot_DB").Collection("prices")
	filter = makeFilter(pricesCollection)
	for i, good := range page.Goods {
		_, err = pricesCollection.UpdateOne(ctx, filter, bson.M{"$set": map[string]float64{good.Custom: prices[i]}})
		if err != nil {
			return err
		}
	}

	// Добавление кодов
	codes := database.MongoDB.Database("play_bot_DB").Collection("codes")
	filter = makeFilter(codes)
	for code, text := range codesTexts {
		_, err = codes.UpdateOne(ctx, filter, bson.M{"$set": map[string]string{code: text}})
		if err != nil {
			return err
		}
	}
	return nil
}

// ================HELPERS=============================HELPERS=============================HELPERS=============
func parseMongoMap[T any](collection *mongo.Collection, protoMap map[string]T) error {
	cur, _ := collection.Find(context.Background(), bson.D{})
	cur.Next(context.Background())
	vals, err := cur.Current.Elements()
	if err != nil {
		return err
	}
	for i, val := range vals {
		if i == 0 {
			continue
		}
		var mapp T
		err := val.Value().Unmarshal(&mapp)
		if err != nil {
			return err
		}
		protoMap[val.Key()] = mapp
	}
	return nil
}

func editPages(pages *mongo.Collection, ctx context.Context, page *structures.TypicalPage, goodPages []*structures.TypicalPage, mainCommandName string) error {
	cur, _ := pages.Find(context.Background(), bson.D{})
	cur.Next(context.Background())
	rawFilter := cur.Current.Lookup("_id")
	filter := bson.D{}
	rawFilter.Unmarshal(&filter)
	prevPageCur := cur.Current.Lookup(page.PrevPage)
	prevPage := structures.TypicalPage{}
	err := prevPageCur.Unmarshal(&prevPage)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return err
	}
	edited := false
	for i, commandRow := range prevPage.Commands {
		for j, commmand := range commandRow {
			if mainCommandName == commmand.Text {
				prevPage.Commands[i][j].Command = page.MainCommand
				edited = true
				break
			}
		}
		if edited {
			break
		}
		if i == len(prevPage.Commands)-1 {
			if len(prevPage.Commands[i]) == 1 {
				prevPage.Commands[i] = append(prevPage.Commands[i],
					structures.Command{Text: mainCommandName, Command: page.MainCommand})
			} else {
				prevPage.Commands = append(prevPage.Commands,
					[]structures.Command{{Text: mainCommandName, Command: page.MainCommand}})
			}

		}
	}
	_, err = pages.UpdateOne(ctx, filter, bson.M{"$set": map[string]*structures.TypicalPage{page.PrevPage: &prevPage}})
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return err
	}
	_, err = pages.UpdateOne(ctx, filter, bson.M{"$set": map[string]*structures.TypicalPage{page.MainCommand: page}})
	if err != nil {
		return err
	}
	if len(goodPages) != 0 {
		for i, goodPage := range goodPages {
			_, err := pages.UpdateOne(ctx, filter, bson.M{"$set": map[string]*structures.TypicalPage{page.Goods[i].Command: goodPage}})
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func makeFilter(collection *mongo.Collection) *bson.D {
	cur, _ := collection.Find(context.Background(), bson.D{})
	cur.Next(context.Background())
	rawFilter := cur.Current.Lookup("_id")
	filter := bson.D{}
	rawFilter.Unmarshal(&filter)
	return &filter
}
