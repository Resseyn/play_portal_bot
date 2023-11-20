package botBase

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"play_portal_bot/internal/botBase/botCommands"
	"play_portal_bot/internal/botBase/botLogic"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/loggers"
)

func BotStart() {
	bot, err := tgbotapi.NewBotAPI(BotKey)
	if err != nil {
		loggers.ErrorLogger.Panic(err)
	}

	bot.Debug = false

	loggers.GlobalLogger.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				fmt.Println(update.Message.MessageID)
				switch update.Message.Command() {
				case "start":
					botCommands.BotStart(bot, &update)
				}
			}
		} else if update.CallbackQuery != nil {
			messageData := helpingMethods.ParseData(update.CallbackQuery.Data)
			fmt.Println(messageData)
			switch messageData.Command {
			case "mainMenu":
				botLogic.Menu(bot, &update, messageData)
			case "showShop":
				botLogic.ShowShop(bot, &update, messageData)
				//case "showPersonalArea":
				//	botCommands.Shop(bot, &update)
				//case "showFAQ":
				//	botCommands.Shop(bot, &update)
				//case "showSupport":
				//	botCommands.Shop(bot, &update)
			}
		}

	}

}
