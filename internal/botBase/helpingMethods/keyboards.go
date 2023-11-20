package helpingMethods

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"play_portal_bot/pkg/utils/structures"
)

// TODO: ПАДРАЧИТЬ ХУЙ

//func CreateInline(data *structures.MessageData, rows, columns int, commands ...structures.Command) *tgbotapi.InlineKeyboardMarkup {
//
//	if len(commands) != rows*columns {
//		panic(fmt.Errorf("ТЫ ЕБАНАТ ПОСЧИТАЙ СТРОЧКИ И СТОЛБЦЫ"))
//	}
//	resrows := make([][]tgbotapi.InlineKeyboardButton, rows)
//	for i := range resrows {
//		resrows[i] = make([]tgbotapi.InlineKeyboardButton, columns)
//	}
//	prev := false
//	if data.PrevCommand != "" {
//		rows++
//		resrows = make([][]tgbotapi.InlineKeyboardButton, rows)
//		for i := range resrows {
//			if i == rows {
//				resrows[i] = make([]tgbotapi.InlineKeyboardButton, 1)
//				break
//			}
//			resrows[i] = make([]tgbotapi.InlineKeyboardButton, columns)
//		}
//		rows--
//		prev = true
//	}
//	cmdcount := 0
//	for row := 0; row < rows; row++ {
//		for column := 0; column < columns; column++ {
//			dataFormat := fmt.Sprintf("%v,%v,%v,%v", data.ChatID, data.MessageID, commands[cmdcount].Command, data.Command)
//
//			resrows[row][column] = tgbotapi.NewInlineKeyboardButtonData(commands[cmdcount].Text, dataFormat)
//
//			cmdcount++
//		}
//	}
//	if prev {
//		backFormat := fmt.Sprintf("%v,%v,%v,%v", data.ChatID, data.MessageID, data.PrevCommand, "")
//		resrows[rows][0] = tgbotapi.NewInlineKeyboardButtonData("Назад", backFormat)
//	}
//
//	kb := tgbotapi.NewInlineKeyboardMarkup(resrows...)
//	return &kb
//}

/*
НОВАЯ КЛАВИАТУРА:
теперь вместо rows и columns используется слайс с количеством элементов в строке
раньше:
	3 rows 2 columns:
---------------------
|	button, button  |
|	button, button  |
|   button, button  |
---------------------
теперь: positions = []int{3,1,2}
----------------------------
|  button, button, button  |
|          button          |
|      button,button       |
----------------------------
*/

func CreateInline(data *structures.MessageData, positions []int, commands ...structures.Command) *tgbotapi.InlineKeyboardMarkup {
	rows := len(positions)
	su := 0
	for _, v := range positions {
		su += v
	}

	if su-len(commands) != 0 {
		panic(fmt.Errorf("ТЫ ЕБАНАТ ПОСЧИТАЙ СТРОЧКИ И СТОЛБЦЫ"))
	}
	resrows := make([][]tgbotapi.InlineKeyboardButton, rows)
	for k, v := range positions {
		resrows[k] = make([]tgbotapi.InlineKeyboardButton, v)
	}
	prev := false
	if data.PrevCommand != "" {
		positions = append(positions, 1)
		rows = len(positions)
		resrows = make([][]tgbotapi.InlineKeyboardButton, rows)
		for k, v := range positions {
			if k == rows {
				resrows[k] = make([]tgbotapi.InlineKeyboardButton, 1)
				break
			}
			resrows[k] = make([]tgbotapi.InlineKeyboardButton, v)
		}
		rows--
		prev = true
	}
	cmdcount := 0
	for row := 0; row < rows; row++ {
		for column := 0; column < positions[row]; column++ {
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
