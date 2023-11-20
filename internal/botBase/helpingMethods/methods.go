package helpingMethods

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
	"strconv"
	"strings"
)

// СОЗДАЕТ НУЖНОЕ СООБЩЕНИЕ
func CreateMessage(chatID int64, picPath, messageContent string, commands *[]structures.Command, messageData *structures.MessageData, positions []int) *tgbotapi.PhotoConfig {
	picBytes, err := ioutil.ReadFile(picPath)
	if err != nil {
		loggers.ErrorLogger.Println(err)
	}
	msg := tgbotapi.NewPhoto(chatID, tgbotapi.FileBytes{Name: "cat1", Bytes: picBytes})
	msg.Caption = messageContent
	msg.ReplyMarkup = CreateInline(messageData, positions,
		*commands...)
	return &msg
}

func ParseData(callbackData string) *structures.MessageData {
	data := strings.Split(callbackData, ",") //0 - chatID 1- messageID 2 - command 3 - prevCommand
	data0, _ := strconv.Atoi(data[0])
	data1, _ := strconv.Atoi(data[0])
	messageData := &structures.MessageData{
		MessageID:   data0,
		ChatID:      int64(data1),
		Command:     data[2],
		PrevCommand: data[3],
	}
	return messageData
}
