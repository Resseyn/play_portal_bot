package botLogic

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"play_portal_bot/pkg/utils/structures"
)

func SendInline(data *structures.MessageData) *tgbotapi.InlineKeyboardMarkup {
	dataFormat := fmt.Sprintf("%v,%v,%v", data.ChatID, data.MessageID, data.Command)
	switch data.Command {
	case "start":
		mainMenuKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Магазин", dataFormat),
				tgbotapi.NewInlineKeyboardButtonData("Кабинет", dataFormat)),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("FAQ", dataFormat),
				tgbotapi.NewInlineKeyboardButtonData("Поддержка", dataFormat)),
		)
		return &mainMenuKeyboard
	case "":

	}
	return nil
}
