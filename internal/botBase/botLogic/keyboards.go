package botLogic

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"play_portal_bot/pkg/utils/structures"
)

// makeInline(&data,[]zalupa{})
//func SendInline(data *structures.MessageData, rows, stolbiki int) *tgbotapi.InlineKeyboardMarkup {
//	dataFormat := fmt.Sprintf("%v,%v,%v", data.ChatID, data.MessageID, data.Command)
//	switch data.Command {
//	case "start":
//		mainMenuKeyboard := tgbotapi.NewInlineKeyboardMarkup(
//			tgbotapi.NewInlineKeyboardRow(
//				tgbotapi.NewInlineKeyboardButtonData("Магазин", dataFormat),
//				tgbotapi.NewInlineKeyboardButtonData("Кабинет", dataFormat)),
//			tgbotapi.NewInlineKeyboardRow(
//				tgbotapi.NewInlineKeyboardButtonData("FAQ", dataFormat),
//				tgbotapi.NewInlineKeyboardButtonData("Поддержка", dataFormat)),
//		)
//		return &mainMenuKeyboard
//	case "":
//
//	}
//	return nil
//}

func SendInline(data *structures.MessageData, rows, columns int, commands ...structures.Command) *tgbotapi.InlineKeyboardMarkup {
	dataFormat := fmt.Sprintf("%v,%v,%v,%v", data.ChatID, data.MessageID, data.Command.Text, data.Command.Command)

	resrows := make([][]tgbotapi.InlineKeyboardButton, rows)
	for i := range resrows {
		resrows[i] = make([]tgbotapi.InlineKeyboardButton, columns)
	}

	cmdcount := 0
	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {

		}
	}

	kb := tgbotapi.NewInlineKeyboardMarkup(resrows...)
	return &kb
}
