package botBase

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"play_portal_bot/internal/botBase/botLogic"
	"play_portal_bot/internal/loggers"
)

func BotStart() {
	bot, err := tgbotapi.NewBotAPI(BotKey)
	if err != nil {
		loggers.ErrorLogger.Panic(err)
	}

	botLogic.DoSomething()

	bot.Debug = false

	loggers.GlobalLogger.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "123")
				_, err := bot.Send(msg)
				if err != nil {
					return
				}
			}
		}

	}

}
