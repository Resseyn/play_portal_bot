package botLogic

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
	prev := false
	if data.PrevCommand != "" {
		rows++
		resrows = make([][]tgbotapi.InlineKeyboardButton, rows)
		for i := range resrows {
			if i == rows {
				resrows[i] = make([]tgbotapi.InlineKeyboardButton, 1)
				break
			}
			resrows[i] = make([]tgbotapi.InlineKeyboardButton, columns)
		}
		rows--
		prev = true
	}

	cmdcount := 0
	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			dataFormat := fmt.Sprintf("%v,%v,%v,%v", data.ChatID, data.MessageID, commands[cmdcount].Command, data.Command)

			resrows[row][column] = tgbotapi.NewInlineKeyboardButtonData(commands[cmdcount].Text, dataFormat)

			cmdcount++
		}
	}
	if prev {
		backFormat := fmt.Sprintf("%v,%v,%v,%v", data.ChatID, data.MessageID, data.PrevCommand, "")
		resrows[rows][0] = tgbotapi.NewInlineKeyboardButtonData("Назад", backFormat)
	}

	kb := tgbotapi.NewInlineKeyboardMarkup(resrows...)
	return &kb
}
