package botLogic

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"play_portal_bot/internal/loggers"
)

func ShowShop(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {

	msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, "pkg/utils/data/img/pngtree-isolated-cat-on-white-background-png-image_7094927.png")
	msg.Caption = "Выберите категорию"
	_, err := bot.Send(msg)
	if err != nil {
		loggers.ErrorLogger.Println(err.Error())
	}
}
