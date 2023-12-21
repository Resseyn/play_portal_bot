package databaseModels

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
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
	cur, _ := pages.Find(context.Background(), bson.D{})
	cur.Next(context.Background())
	vals, err := cur.Current.Elements()
	futurePages := make(map[string]*structures.TypicalPage)
	for i, val := range vals {
		if i == 0 {
			continue
		}
		mapp := structures.TypicalPage{}
		val.Value().Unmarshal(&mapp)
		futurePages[val.Key()] = &mapp
	}
	structures.Pages = futurePages
	if err != nil {
		return err
	}
	cur, _ = handlers.Find(context.Background(), bson.D{})
	cur.Next(context.Background())
	vals, err = cur.Current.Elements()
	futureHandlers := make(map[string]string)
	for i, val := range vals {
		if i == 0 {
			continue
		}
		var mapp string
		val.Value().Unmarshal(&mapp)
		futureHandlers[val.Key()] = mapp
	}
	structures.Handlers = futureHandlers
	if err != nil {
		return err
	}
	cur, _ = pages.Find(context.Background(), bson.D{})
	cur.Next(context.Background())
	vals, err = cur.Current.Elements()
	futurePages = make(map[string]*structures.TypicalPage)
	for i, val := range vals {
		if i == 0 {
			continue
		}
		mapp := structures.TypicalPage{}
		val.Value().Unmarshal(&mapp)
		futurePages[val.Key()] = &mapp
	}
	structures.Pages = futurePages
	if err != nil {
		return err
	}
	cur, _ = parameters.Find(context.Background(), bson.D{})
	cur.Next(context.Background())
	vals, err = cur.Current.Elements()
	futureParameters := make(map[string][]string)
	for i, val := range vals {
		if i == 0 {
			continue
		}
		mapp := make([]string, 10)
		val.Value().Unmarshal(&mapp)
		futureParameters[val.Key()] = mapp
	}
	structures.Parameters = futureParameters
	if err != nil {
		return err
	}
	cur, _ = commands.Find(context.Background(), bson.D{})
	cur.Next(context.Background())
	vals, err = cur.Current.Elements()
	futureCommands := make(map[string]string)
	for i, val := range vals {
		if i == 0 {
			continue
		}
		var mapp string
		val.Value().Unmarshal(&mapp)
		futureCommands[val.Key()] = mapp
	}
	structures.Commands = futureCommands
	if err != nil {
		return err
	}
	cur, _ = prices.Find(context.Background(), bson.D{})
	cur.Next(context.Background())
	vals, err = cur.Current.Elements()
	futurePrices := make(map[string]float64)
	for i, val := range vals {
		if i == 0 {
			continue
		}
		var mapp float64
		val.Value().Unmarshal(&mapp)
		futurePrices[val.Key()] = mapp
	}
	structures.Prices = futurePrices
	if err != nil {
		return err
	}
	cur, _ = codes.Find(context.Background(), bson.D{})
	cur.Next(context.Background())
	vals, err = cur.Current.Elements()
	futureCodes := make(map[string]string)
	for i, val := range vals {
		if i == 0 {
			continue
		}
		var mapp string
		val.Value().Unmarshal(&mapp)
		futureCodes[val.Key()] = mapp
	}
	structures.Codes = futureCodes
	if err != nil {
		return err
	}
	return nil
}

func AddNewPageToMongo(page *structures.TypicalPage, handlerName string, handlerParams []string, prices []float64, codesTexts map[string]string, goodsNames []string) error {
	////TODO: тестовые значения
	//goodsNames = []string{"tovar1,", "tovar2"}
	//page = &structures.TypicalPage{
	//	URL:         "pkg/utils/data/img/shopImages/servicesImages/twitch.png",
	//	Text:        "cmmfafasfas",
	//	MainCommand: helpingMethods.RandStringRunes(4), //TODO: при заполнении всегда проверять, нет ли такой комманды в словаре Commands
	//	Commands: [][]structures.Command{
	//		{
	//			{Text: goodsNames[0], Command: helpingMethods.RandStringRunes(4)},
	//			{Text: goodsNames[1], Command: helpingMethods.RandStringRunes(4)},
	//		},
	//	},
	//	PrevPage: structures.Commands["mainMenu"],
	//}
	//page.Goods = []structures.Good{
	//	{URL: "pkg/utils/data/img/shopImages/servicesImages/spotify/spotify_individual_1.jpg",
	//		Text:    "Выберите товар",
	//		Custom:  helpingMethods.RandStringRunes(4),
	//		Command: page.Commands[0][0].Command},
	//	{URL: "pkg/utils/data/img/shopImages/servicesImages/spotify/spotify_individual_3.jpg",
	//		Text:    "Выберите товар",
	//		Custom:  helpingMethods.RandStringRunes(4),
	//		Command: page.Commands[0][1].Command}, //TODO:удачи автоматом это все правильно распихивать, 100500 строк генераторов
	//}
	//handlerName = "newSpotHandler"
	//handlerParams = []string{"Введите", "dssdfsd"}
	//prices = []float64{333.0, 666.0}
	//codesTexts = map[string]string{page.MainCommand: "Новый спотик"}
	//codesTexts - словарь с кастомом товара и его названием, которое юзер должен ввести в хендлере до этой функции
	for i, good := range page.Goods {
		codesTexts[good.Custom] = goodsNames[i]
	}
	ctx := context.Background()
	// Добавление страницы
	pages := database.MongoDB.Database("play_bot_DB").Collection("pages")
	cur, _ := pages.Find(context.Background(), bson.D{})
	cur.Next(context.Background())
	rawFilter := cur.Current.Lookup("_id")
	filter := bson.D{}
	rawFilter.Unmarshal(&filter)
	_, err := pages.UpdateOne(ctx, filter, bson.M{"$set": map[string]*structures.TypicalPage{page.MainCommand: page}})
	if err != nil {
		return err
	}

	// Добавление обработчика
	handlers := database.MongoDB.Database("play_bot_DB").Collection("handlers")
	cur, _ = handlers.Find(context.Background(), bson.D{})
	cur.Next(context.Background())
	rawFilter = cur.Current.Lookup("_id")
	filter = bson.D{}
	rawFilter.Unmarshal(&filter)
	for _, good := range page.Goods {
		_, err = handlers.UpdateOne(ctx, filter, bson.M{"$set": map[string]string{good.Custom: handlerName}})
		if err != nil {
			return err
		}
	}

	// Добавление параметров
	parameters := database.MongoDB.Database("play_bot_DB").Collection("parameters")
	cur, _ = parameters.Find(context.Background(), bson.D{})
	cur.Next(context.Background())
	rawFilter = cur.Current.Lookup("_id")
	filter = bson.D{}
	rawFilter.Unmarshal(&filter)
	_, err = parameters.UpdateOne(ctx, filter, bson.M{"$set": map[string][]string{handlerName: handlerParams}})
	if err != nil {
		return err
	}

	// Добавление команд
	commandsColl := database.MongoDB.Database("play_bot_DB").Collection("commands")
	cur, _ = commandsColl.Find(context.Background(), bson.D{})
	cur.Next(context.Background())
	rawFilter = cur.Current.Lookup("_id")
	filter = bson.D{}
	rawFilter.Unmarshal(&filter)
	for _, commands := range page.Commands {
		for _, command := range commands {
			_, err = commandsColl.UpdateOne(ctx, filter, bson.M{"$set": map[string]string{command.Text: command.Command}})
			if err != nil {
				return err
			}
		}

	}

	// Добавление цен
	pricesCollection := database.MongoDB.Database("play_bot_DB").Collection("prices")
	cur, _ = pricesCollection.Find(context.Background(), bson.D{})
	cur.Next(context.Background())
	rawFilter = cur.Current.Lookup("_id")
	filter = bson.D{}
	rawFilter.Unmarshal(&filter)
	for i, good := range page.Goods {
		_, err = pricesCollection.UpdateOne(ctx, filter, bson.M{"$set": map[string]float64{good.Custom: prices[i]}})
		if err != nil {
			return err
		}
	}

	// Добавление кодов
	codes := database.MongoDB.Database("play_bot_DB").Collection("codes")
	cur, _ = codes.Find(context.Background(), bson.D{})
	cur.Next(context.Background())
	rawFilter = cur.Current.Lookup("_id")
	filter = bson.D{}
	rawFilter.Unmarshal(&filter)
	for code, text := range codesTexts {
		_, err = codes.UpdateOne(ctx, filter, bson.M{"$set": map[string]string{code: text}})
		if err != nil {
			return err
		}
	}

	return nil
}

//	func UpdateMongo() error {
//		pages := database.MongoDB.Database("play_bot_DB").Collection("pages")
//		handlers := database.MongoDB.Database("play_bot_DB").Collection("handlers")
//		parameters := database.MongoDB.Database("play_bot_DB").Collection("parameters")
//		commands := database.MongoDB.Database("play_bot_DB").Collection("commands")
//		prices := database.MongoDB.Database("play_bot_DB").Collection("prices")
//		codes := database.MongoDB.Database("play_bot_DB").Collection("codes")
//		cur, _ := pages.Find(context.Background(), bson.D{})
//		err := cur.Decode(&structures.Pages)
//		if err != nil {
//			return err
//		}
//	}
