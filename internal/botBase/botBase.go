package botBase

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"play_portal_bot/internal/botBase/botCommands"
	"play_portal_bot/internal/loggers"
	"strings"
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

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "start":
					botCommands.BotStart(bot, &update)
				}
			}
		} else if update.CallbackQuery != nil {
			data := strings.Split(update.CallbackQuery.Data, ",") //0 - chatID 1- messageID 2 - command 3 - prevCommand
			switch data[2] {
			case "mainMenu":
				botCommands.Menu(bot, &update)
			case "showShop":
				botCommands.Shop(bot, &update)
			case "showPersonalArea":
				botCommands.Shop(bot, &update)
			case "showFAQ":
				botCommands.Shop(bot, &update)
			case "showSupport":
				botCommands.Shop(bot, &update)
			}
		}

	}

}
