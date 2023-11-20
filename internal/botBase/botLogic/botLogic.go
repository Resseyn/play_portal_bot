package botLogic

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
)

func Menu(bot *tgbotapi.BotAPI, update *tgbotapi.Update, data *structures.MessageData) {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/Hydra.webp"
	messageContent := "МАГАЗИН ИГР 'ГИДРА'"
	data.Command = "mainMenu"
	data.PrevCommand = ""
	commands := &[]structures.Command{
		{Text: "Магазин", Command: "showShop"},
		{Text: "Кабинет", Command: "showPersonalArea"},
		{Text: "Поддержка", Command: "showSupport"},
		{Text: "FAQ", Command: "showFAQ"},
	}
	positions := []int{2, 2}
	data.PrevCommand = ""
	// =========PARAMS=========

	editConf := helpingMethods.EditMessageWithPhotoAndReplyMarkup(data, commands, messageContent, picPath, positions)
	_, err := bot.Send(editConf)
	if err != nil {
		loggers.ErrorLogger.Println(err.Error(), "mediaError")
	}
}

func ShowShop(bot *tgbotapi.BotAPI, update *tgbotapi.Update, data *structures.MessageData) {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/gettyimages-1067956982.jpg"
	messageContent := "МАГАЗИН ИГР 'ГИДРА'"
	data.Command = "showShop"
	positions := []int{2, 1}
	commands := &[]structures.Command{
		{Text: "Игровые сервисы", Command: ""},
		{Text: "Сервисы", Command: ""},
		{Text: "НАРКОТА", Command: ""},
	}
	// =========PARAMS=========

	editConf := helpingMethods.EditMessageWithPhotoAndReplyMarkup(data, commands, messageContent, picPath, positions)
	_, err := bot.Send(editConf)
	if err != nil {
		loggers.ErrorLogger.Println(err.Error(), "mediaError")
	}
}
