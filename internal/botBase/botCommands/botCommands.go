package botCommands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"play_portal_bot/internal/botBase/botLogic"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
)

func BotStart(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, "pkg/utils/data/img/pngtree-isolated-cat-on-white-background-png-image_7094927.png")
	msg.Caption = "Описание изображения"
	msg.ReplyMarkup = botLogic.CreateInline(
		&structures.MessageData{
			MessageID:   update.Message.MessageID,
			ChatID:      update.Message.Chat.ID,
			Command:     "start",
			PrevCommand: "",
		}, 2, 2,
		[]structures.Command{
			{
				Text:    "Магазин",
				Command: "showShop",
			},
			{
				Text:    "Кабинет",
				Command: "showPersonalArea",
			},
			{
				Text:    "Поддержка",
				Command: "showSupport",
			},
			{
				Text:    "FAQ",
				Command: "showFAQ",
			},
		}...)
	_, err := bot.Send(msg)
	if err != nil {
		loggers.ErrorLogger.Println(err.Error())
	}
}
