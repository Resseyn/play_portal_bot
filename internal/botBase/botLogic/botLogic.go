package botLogic

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
)

func Menu(bot *tgbotapi.BotAPI, update *tgbotapi.Update, data *structures.MessageData) {

	// =========PARAMS=========
	chatID := update.CallbackQuery.Message.Chat.ID
	picPath := "pkg/utils/data/img/gettyimages-1067956982.jpg"
	messageContent := "МАГАЗИН ИГР 'ГИДРА'"
	commands := &[]structures.Command{
		{Text: "Магазин", Command: "showShop"},
		{Text: "Кабинет", Command: "showPersonalArea"},
		{Text: "Поддержка", Command: "showSupport"},
		{Text: "FAQ", Command: "showFAQ"},
	}
	// =========PARAMS=========

	picBytes, err := ioutil.ReadFile(picPath)
	if err != nil {
		loggers.ErrorLogger.Println(err)
	}
	photo := tgbotapi.NewPhoto(chatID, tgbotapi.FileBytes{Name: "cat2", Bytes: picBytes})
	kb := helpingMethods.CreateInline(data, []int{2, 2}, *commands...)
	editMediaConf := tgbotapi.EditMessageMediaConfig{Media: tgbotapi.InputMediaPhoto{BaseInputMedia: tgbotapi.BaseInputMedia{
		Type:      "photo",
		Media:     photo.File,
		Caption:   messageContent,
		ParseMode: "Markdown",
	},
	},
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:          chatID,
			MessageID:       data.MessageID,
			InlineMessageID: "",
			ReplyMarkup:     kb,
		}}
	_, err = bot.Send(editMediaConf)
	if err != nil {
		loggers.ErrorLogger.Println(err.Error(), "mediaError")
	}
}

func ShowShop(bot *tgbotapi.BotAPI, update *tgbotapi.Update, data *structures.MessageData) {

	// =========PARAMS=========
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
	// =========PARAMS=========

	msg := helpingMethods.CreateMessage(chatID, picPath, messageContent, commands, messageData, positions)
	_, err := bot.Send(msg)
	if err != nil {
		loggers.ErrorLogger.Println(err.Error())
	}
}
