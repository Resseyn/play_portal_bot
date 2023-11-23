package helpingMethods

import (
	"play_portal_bot/pkg/utils/structures"
	"strconv"
	"strings"
)

// import (
//
//	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
//	"gopkg.in/telebot.v3"
//	"io/ioutil"
//	"play_portal_bot/internal/loggers"
//	"play_portal_bot/pkg/utils/structures"
//	"strconv"
//	"strings"
//
// )
//
// // СОЗДАЕТ НУЖНОЕ СООБЩЕНИЕ
//
//	func CreateMessage(chatID int64, picPath, messageContent string, commands *[]structures.Command, messageData *structures.MessageData, positions []int) *telebot.Message {
//		//picBytes, err := ioutil.ReadFile(picPath)
//		//if err != nil {
//		//	loggers.ErrorLogger.Println(err)
//		//}
//		//msg := tgbotapi.NewPhoto(chatID, tgbotapi.FileBytes{Name: "cat1", Bytes: picBytes})
//		//msg.Caption = messageContent
//		//msg.ReplyMarkup = CreateInline(messageData, positions,
//		//	*commands...)
//		//return &msg
//
//		msg := telebot.Message{Photo: }
//		msg.Caption = messageContent
//		msg.ReplyMarkup = CreateInline(messageData, positions,
//			*commands...)
//		return &msg
//	}

// ParseData парсит строку с колбек кновпеи в структуру
func ParseData(callbackData string) *structures.MessageData {
	data := strings.Split(callbackData, ",") //0 - chatID 1- messageID 2 - command 3 - prevCommand
	data0, _ := strconv.Atoi(data[0][1:])    //очень опасная хуйня какая-то залупа первым символом попарает весь атой ломает
	data1, _ := strconv.Atoi(data[1])
	messageData := &structures.MessageData{
		ChatID:      int64(data0),
		MessageID:   data1,
		Command:     data[2],
		PrevCommand: data[3],
	}
	return messageData
}

//
//// EditMessageWithPhotoAndReplyMarkup ИЗМЕНЯЕТ СООБЩЕНИЕ С КАРТИНКОЙ И КЛАВОЙ
//func EditMessageWithPhotoAndReplyMarkup(data *structures.MessageData, commands *[]structures.Command, messageContent, picPath string, positions []int) *tgbotapi.EditMessageMediaConfig {
//	picBytes, err := ioutil.ReadFile(picPath)
//	if err != nil {
//		loggers.ErrorLogger.Println(err)
//	}
//	photo := tgbotapi.NewPhoto(data.ChatID, tgbotapi.FileBytes{Name: "cat2", Bytes: picBytes})
//	kb := CreateInline(data, positions, *commands...)
//	editMediaConf := &tgbotapi.EditMessageMediaConfig{Media: tgbotapi.InputMediaPhoto{BaseInputMedia: tgbotapi.BaseInputMedia{
//		Type:      "photo",
//		Media:     photo.File,
//		Caption:   messageContent,
//		ParseMode: "Markdown",
//	},
//	},
//		BaseEdit: tgbotapi.BaseEdit{
//			ChatID:          data.ChatID,
//			MessageID:       data.MessageID,
//			InlineMessageID: "",
//			ReplyMarkup:     kb,
//		}}
//	return editMediaConf
//}
