package databaseModels

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/pkg/database"
	"play_portal_bot/pkg/utils/structures"
)

func GetAllOfMapsFromMongo() error {
	//pages := database.MongoDB.Database("play_bot_DB").Collection("pages")
	//handlers := database.MongoDB.Database("play_bot_DB").Collection("handlers")
	//parameters := database.MongoDB.Database("play_bot_DB").Collection("parameters")
	//commands := database.MongoDB.Database("play_bot_DB").Collection("commands")
	//prices := database.MongoDB.Database("play_bot_DB").Collection("prices")
	//codes := database.MongoDB.Database("play_bot_DB").Collection("codes")
	//cur, _ := pages.Find(context.Background(), bson.D{})
	//err := cur.Decode(&structures.Pages)
	//if err != nil {
	//	return err
	//} //TODO: ну вот короче так все из монго достаешь и запихиваешь по словарям
	return nil
}

func AddNewPageToMongo(page *structures.TypicalPage, handlerName string, handlerParams []string, prices []float64, codesTexts map[string]string, goodsNames []string) error {
	//TODO: тестовые значения
	goodsNames = []string{"tovar1,", "tovar2"}
	page = &structures.TypicalPage{
		URL:         "pkg/utils/data/img/shopImages/servicesImages/twitch.png",
		Text:        "cmmfafasfas",
		MainCommand: helpingMethods.RandStringRunes(4), //TODO: при заполнении всегда проверять, нет ли такой комманды в словаре Commands
		Commands: [][]structures.Command{
			{
				{Text: goodsNames[0], Command: helpingMethods.RandStringRunes(4)},
				{Text: goodsNames[1], Command: helpingMethods.RandStringRunes(4)},
			},
		},
		PrevPage: structures.Commands["mainMenu"],
	}
	page.Goods = []structures.Good{
		{URL: "pkg/utils/data/img/shopImages/servicesImages/spotify/spotify_individual_1.jpg",
			Text:    "Выберите товар",
			Custom:  helpingMethods.RandStringRunes(4),
			Command: page.Commands[0][0].Command},
		{URL: "pkg/utils/data/img/shopImages/servicesImages/spotify/spotify_individual_3.jpg",
			Text:    "Выберите товар",
			Custom:  helpingMethods.RandStringRunes(4),
			Command: page.Commands[0][1].Command}, //TODO:удачи автоматом это все правильно распихивать, 100500 строк генераторов
	}
	handlerName = "newSpotHandler"
	handlerParams = []string{"Введите", "dssdfsd"}
	prices = []float64{333.0, 666.0}
	codesTexts = map[string]string{page.MainCommand: "Новый спотик"}
	//codesTexts - словарь с кастомом товара и его названием, которое юзер должен ввести в хендлере до этой функции
	//TODO: там где циклы создать заранее нормальный словарь, уже его потом добавлять в монгу, чтобы потом достать целый словарь и его юзать
	for i, good := range page.Goods {
		codesTexts[good.Custom] = goodsNames[i]
	}
	ctx := context.Background()
	// Добавление страницы
	pages := database.MongoDB.Database("play_bot_DB").Collection("pages")
	_, err := pages.InsertOne(ctx, bson.M{page.MainCommand: page})
	if err != nil {
		return err
	}

	// Добавление обработчика
	handlers := database.MongoDB.Database("play_bot_DB").Collection("handlers")
	for _, good := range page.Goods {
		_, err = handlers.InsertOne(ctx, bson.M{good.Custom: handlerName})
		if err != nil {
			return err
		}
	}

	// Добавление параметров
	parameters := database.MongoDB.Database("play_bot_DB").Collection("parameters")
	_, err = parameters.InsertOne(ctx, bson.M{handlerName: handlerParams})
	if err != nil {
		return err
	}

	// Добавление команд
	commandsColl := database.MongoDB.Database("play_bot_DB").Collection("commands")
	for _, commands := range page.Commands {
		for _, command := range commands {
			_, err = commandsColl.InsertOne(ctx, bson.M{command.Text: command.Command})
			if err != nil {
				return err
			}
		}

	}

	// Добавление цен
	pricesCollection := database.MongoDB.Database("play_bot_DB").Collection("prices")
	for i, good := range page.Goods {
		_, err = pricesCollection.InsertOne(ctx, bson.M{good.Custom: prices[i]})
		if err != nil {
			return err
		}
	}

	// Добавление кодов
	codes := database.MongoDB.Database("play_bot_DB").Collection("codes")
	for code, text := range codesTexts {
		_, err = codes.InsertOne(ctx, bson.M{code: text})
		if err != nil {
			return err
		}
	}

	return nil
}
