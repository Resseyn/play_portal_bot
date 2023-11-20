package botBase

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"play_portal_bot/internal/botBase/botCommands"
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

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "start":
					botCommands.BotStart(bot, &update)
				}
			}
		}
		if update.CallbackQuery != nil {
			fmt.Println(update.CallbackQuery.Data)
		}

	}

}
