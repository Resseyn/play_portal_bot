package helpingMethods

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"play_portal_bot/pkg/utils/structures"
)

// CreateInline создает инлайн-клавиатуру с кнопками и кнопкой возврата.
// Каждый вложенный массив команд создает новую строку кнопок.
func CreateInline(data *structures.MessageData, commands ...[]structures.Command) *telebot.ReplyMarkup {
	var rows [][]telebot.InlineButton

	for _, cmdRow := range commands {

		var row []telebot.InlineButton

		for _, cmd := range cmdRow {

			dataFormat := fmt.Sprintf("%v,%v,%v,%v", cmd.Command, data.Command, data.Price, data.Custom)
			button := telebot.InlineButton{
				Data: dataFormat,
				Text: cmd.Text,
			}
			row = append(row, button)
		}
		rows = append(rows, row)
	}
	if data.PrevCommand != "" {
		backFormat := fmt.Sprintf("%v,%v,%v,%v", data.PrevCommand, "", 0, "")

		backButton := telebot.InlineButton{
			Data: backFormat,
			Text: "Назад",
		}
		rows = append(rows, []telebot.InlineButton{backButton})
	}

	kb := &telebot.ReplyMarkup{
		InlineKeyboard: rows,
	}
	return kb
}
