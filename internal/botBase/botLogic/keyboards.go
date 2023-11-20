package botLogic

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
)

type MessageData struct {
	ChatID  int
	Command string
}

func SendInline(userid int64, command string) tgbotapi.InlineKeyboardMarkup {
	msgData = strings.Split("123, start", ",")
	id, _ := strconv.Atoi(msgData[0])
	newData := MessageData{chatid: id, command: msgData[1]}
	var kb = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("MessageData")))
}
