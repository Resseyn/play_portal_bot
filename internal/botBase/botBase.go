package botBase

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"play_portal_bot/internal/loggers"
)

func BotStart() {
	bot, err := tgbotapi.NewBotAPI("6978330124:AAHEaBkrAr4teJ5Fhwbn9skGrQuT6Z_BhX0")
	if err != nil {
		loggers.ErrorLogger.Panic(err)
	}

	bot.Debug = true

	loggers.GlobalLogger.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		fmt.Println(update.Message.Text)
	}

}
