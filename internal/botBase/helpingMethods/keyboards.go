package helpingMethods

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"play_portal_bot/pkg/utils/structures"
)

// TODO: ПАДРАЧИТЬ ХУЙ

// makeInline(&data,[]zalupa{})
//func CreateInline(data *structures.MessageData) *tgbotapi.InlineKeyboardMarkup {
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

func CreateInline(data *structures.MessageData, rows, columns int, commands ...structures.Command) *tgbotapi.InlineKeyboardMarkup {

	if len(commands) != rows*columns {
		panic(fmt.Errorf("ТЫ ЕБАНАТ ПОСЧИТАЙ СТРОЧКИ И СТОЛБЦЫ"))
	}

	resrows := make([][]tgbotapi.InlineKeyboardButton, rows)
	for i := range resrows {
		resrows[i] = make([]tgbotapi.InlineKeyboardButton, columns)
	}

	cmdcount := 0
	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			dataFormat := fmt.Sprintf("%v,%v,%v,%v", data.ChatID, data.MessageID, commands[cmdcount].Command, data.Command)

			resrows[row][column] = tgbotapi.NewInlineKeyboardButtonData(commands[cmdcount].Text, dataFormat)

			cmdcount++
		}
	}

	kb := tgbotapi.NewInlineKeyboardMarkup(resrows...)
	return &kb
}
