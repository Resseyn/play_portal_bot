package helpingMethods

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"play_portal_bot/pkg/utils/structures"
)

// СОЗДАЕТ НУЖНОЕ СООБЩЕНИЕ
func CreateMessage(chatID int64, picPath, messageContent string, commands *[]structures.Command, messageData *structures.MessageData, rows, columns int) *tgbotapi.PhotoConfig {
	msg := tgbotapi.NewPhotoUpload(chatID, picPath)
	msg.Caption = messageContent
	msg.ReplyMarkup = CreateInline(messageData, rows, columns,
		*commands...)
	return &msg
}
