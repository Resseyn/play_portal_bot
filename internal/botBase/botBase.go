package botBase

import (
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/botBase/botCommands"
	"play_portal_bot/internal/loggers"
	"time"
)

func BotStart() {
	//bot, err := tgbotapi.NewBotAPI(BotKey)
	//if err != nil {
	//	loggers.ErrorLogger.Panic(err)
	//}
	//
	//bot.Debug = false
	//
	//loggers.GlobalLogger.Printf("Authorized on account %s", bot.Self.UserName)
	//
	//u := tgbotapi.NewUpdate(0)
	//u.Timeout = 60
	//
	//updates := bot.GetUpdatesChan(u)
	//
	//for update := range updates {
	//	if update.Message != nil {
	//		if update.Message.IsCommand() {
	//			fmt.Println(update.Message.MessageID)
	//			switch update.Message.Command() {
	//			case "start":
	//				botCommands.Start(bot, &update)
	//			}
	//		}
	//	} else if update.CallbackQuery != nil {
	//		messageData := helpingMethods.ParseData(update.CallbackQuery.Data)
	//		fmt.Println(messageData)
	//		switch messageData.Command {
	//		case "mainMenu":
	//			botLogic.Menu(bot, &update, messageData)
	//		case "showShop":
	//			botLogic.ShowShop(bot, &update, messageData)
	//			//case "showPersonalArea":
	//			//	botCommands.Shop(bot, &update)
	//			//case "showFAQ":
	//			//	botCommands.Shop(bot, &update)
	//			//case "showSupport":
	//			//	botCommands.Shop(bot, &update)
	//		}
	//	}
	//
	//}
	settings := telebot.Settings{
		Token:  BotKey,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(settings)
	if err != nil {
		loggers.ErrorLogger.Fatal(err)
		return
	}

	b.Handle("/start", func(c telebot.Context) error {
		return botCommands.Start(c)
	})
	if err != nil {
		loggers.ErrorLogger.Fatal(err)
		panic(err)
	}

	b.Start()
}
