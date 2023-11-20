package botLogic

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
)

func Menu(bot *tgbotapi.BotAPI, update *tgbotapi.Update, data *structures.MessageData) {
	chatID := update.CallbackQuery.Message.Chat.ID
	picPath := "pkg/utils/data/img/gettyimages-1067956982.jpg"
	messageContent := "МАГАЗИН ГИДРА"
	//messageData := &structures.MessageData{
	//	MessageID:   update.Message.MessageID,
	//	ChatID:      update.Message.Chat.ID,
	//	Command:     "start",
	//	PrevCommand: "mainMenu",
	//}
	//rows := 2
	//columns := 2
	commands := &[]structures.Command{
		{Text: "Магазин", Command: "showShop"},
		{Text: "Кабинет", Command: "showPersonalArea"},
		{Text: "Поддержка", Command: "showSupport"},
		{Text: "FAQ", Command: "showFAQ"},
	}
	picBytes, err := ioutil.ReadFile(picPath)
	if err != nil {
		loggers.ErrorLogger.Println(err)
	}
	editMediaConf := tgbotapi.EditMessageMediaConfig{Media: tgbotapi.FileBytes{Name: "cat2", Bytes: picBytes}}
	kb := helpingMethods.CreateInline(data, []int{2, 2}, *commands...)
	editTextConfig := tgbotapi.NewEditMessageTextAndMarkup(chatID, data.MessageID, messageContent, *kb)
	_, err = bot.Send(editTextConfig)
	if err != nil {
		loggers.ErrorLogger.Println(err.Error(), "textAndKbError")
	}
	_, err = bot.Send(editMediaConf)
	if err != nil {
		loggers.ErrorLogger.Println(err.Error(), "mediaError")
	}
}

func ShowShop(bot *tgbotapi.BotAPI, update *tgbotapi.Update, data *structures.MessageData) {
	chatID := update.Message.Chat.ID
	picPath := "pkg/utils/data/img/gettyimages-1067956982.jpg"
	messageContent := "МАГАЗИН ИГР 'ГИДРА'"
	messageData := &structures.MessageData{
		MessageID:   update.Message.MessageID,
		ChatID:      update.Message.Chat.ID,
		Command:     "start",
		PrevCommand: "mainMenu",
	}
	positions := []int{2, 2}
	commands := &[]structures.Command{
		{Text: "Магазин", Command: "showShop"},
		{Text: "Кабинет", Command: "showPersonalArea"},
		{Text: "Поддержка", Command: "showSupport"},
		{Text: "FAQ", Command: "showFAQ"},
	}
	msg := helpingMethods.CreateMessage(chatID, picPath, messageContent, commands, messageData, positions)
	_, err := bot.Send(msg)
	if err != nil {
		loggers.ErrorLogger.Println(err.Error())
	}
}
