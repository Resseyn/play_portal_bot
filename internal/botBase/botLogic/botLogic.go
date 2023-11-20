package botLogic

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
)

func Menu(bot *tgbotapi.BotAPI, update *tgbotapi.Update, data *structures.MessageData) {
	//// Create a FileBytes from the temporary file
	//fileBytes := tgbotapi.FileBytes{
	//	Name:  fileName,
	//	Bytes: imageBytes,
	//}
	//
	//baseInputMedia := tgbotapi.BaseInputMedia{
	//	Type:      "photo", // Set the desired media type
	//	Media:     fileBytes,
	//	ParseMode: "markdown", // Set the desired parse mode
	//}
	//
	//// Create an EditMessageMediaConfig to update the message
	//editMessageConfig := tgbotapi.EditMessageMediaConfig{
	//	BaseEdit: tgbotapi.BaseEdit{
	//		ChatID:    chatID,
	//		MessageID: messageID,
	//	},
	//	Media: tgbotapi.InputMediaPhoto{
	//		BaseInputMedia: baseInputMedia,
	//	},
	//}
	//
	//// Edit the existing message with the updated photo and inline keyboard
	//_, err = bot.Send(&editMessageConfig)
	//if err != nil {
	//	return fmt.Errorf("error editing message: %v", err)
	//}
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
	editMediaConf := tgbotapi.EditMessageMediaConfig{Media: tgbotapi.InputMediaPhoto{BaseInputMedia: tgbotapi.BaseInputMedia{
		Type:      "photo",
		Media:     photo.File,
		Caption:   "Обновленная картинка",
		ParseMode: "Markdown",
	},
	},
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:          chatID,
			MessageID:       data.MessageID,
			InlineMessageID: "",
		}}
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
