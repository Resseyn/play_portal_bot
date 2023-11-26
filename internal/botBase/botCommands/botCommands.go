package botCommands

import (
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/pkg/utils/structures"
)

// Start КОМАНДА СТАРТА, ПО СОВМЕСТИТЕЛЬСТВУ ВЫВОД ГЛАВНОГО МЕНЮ
func Start(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/mainMenuImages/Hydra.webp"
	messageContent := "МАГАЗИН ИГР 'ГИДРА'"
	messageData := &structures.MessageData{
		MessageID:   c.Message().ID,
		ChatID:      c.Chat().ID,
		Command:     structures.Commands["mainMenu"],
		PrevCommand: "",
		Price:       0,
	}
	commands := [][]structures.Command{
		{
			{Text: "Магазин", Command: structures.Commands["shop"]},
			{Text: "Кабинет", Command: structures.Commands["personalCabinet"]}},
		{
			{Text: "Поддержка", Command: structures.Commands["support"]},
			{Text: "FAQ", Command: structures.Commands["faq"]}},
	}
	// =========PARAMS=========

	msg := &telebot.Photo{
		File:    telebot.FromDisk(picPath),
		Caption: messageContent,
	}
	keyboard := helpingMethods.CreateInline(messageData, commands...)
	err := c.Send(msg, &telebot.SendOptions{
		ParseMode:   telebot.ModeHTML,
		ReplyMarkup: keyboard,
	})
	return err

}

// /tool1 - команда для вызова админ панели
// ваще на самом деле я бы отдельного бота создал который бы присылал нам че то
// ну допустим жоские уведы про покупки или там сообщения о багах но похуй пусть пока будет мне захотелось написать
func CreateAdminPanel(c telebot.Context) error {

	picPath := "pkg/utils/data/img/adminImages/ramzes.jpg"
	messageContent := "админ менюшка"
	messageData := &structures.MessageData{
		MessageID:   c.Message().ID,
		ChatID:      c.Chat().ID,
		Command:     structures.Commands["adminPanel"],
		PrevCommand: "",
		Price:       0,
	}
	commands := [][]structures.Command{
		{
			{Text: "Показать репорты", Command: structures.Commands["showReports"]},
			{Text: "Падрачить хуй", Command: structures.Commands["showAdminPanel"]}}}

	msg := &telebot.Photo{
		File:    telebot.FromDisk(picPath),
		Caption: messageContent,
	}
	keyboard := helpingMethods.CreateInline(messageData, commands...)
	err := c.Send(msg, &telebot.SendOptions{
		ParseMode:   telebot.ModeHTML,
		ReplyMarkup: keyboard,
	})

	return err

}
