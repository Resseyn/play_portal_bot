package botLogic

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
)

func Menu(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	chatID := update.CallbackQuery.Message.Chat.ID
	picPath := "pkg/utils/data/img/gettyimages-1067956982.jpg"
	messageContent := "МАГАЗИН ИГР 'ГИДРА'"
	messageData := &structures.MessageData{
		MessageID:   update.Message.MessageID,
		ChatID:      update.Message.Chat.ID,
		Command:     "start",
		PrevCommand: "mainMenu",
	}
	rows := 2
	columns := 2
	commands := &[]structures.Command{
		{Text: "Магазин", Command: "showShop"},
		{Text: "Кабинет", Command: "showPersonalArea"},
		{Text: "Поддержка", Command: "showSupport"},
		{Text: "FAQ", Command: "showFAQ"},
	}
	msg := helpingMethods.CreateMessage(chatID, picPath, messageContent, commands, messageData, rows, columns)
	_, err := bot.Edit(msg)
	if err != nil {
		loggers.ErrorLogger.Println(err.Error())
	}
}

func ShowShop(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	picPath := "pkg/utils/data/img/gettyimages-1067956982.jpg"
	messageContent := "МАГАЗИН ИГР 'ГИДРА'"
	messageData := &structures.MessageData{
		MessageID:   update.Message.MessageID,
		ChatID:      update.Message.Chat.ID,
		Command:     "start",
		PrevCommand: "mainMenu",
	}
	rows := 2
	columns := 2
	commands := &[]structures.Command{
		{Text: "Магазин", Command: "showShop"},
		{Text: "Кабинет", Command: "showPersonalArea"},
		{Text: "Поддержка", Command: "showSupport"},
		{Text: "FAQ", Command: "showFAQ"},
	}
	msg := helpingMethods.CreateMessage(chatID, picPath, messageContent, commands, messageData, rows, columns)
	_, err := bot.Edit(msg)
	if err != nil {
		loggers.ErrorLogger.Println(err.Error())
	}
}
