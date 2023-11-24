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
		Command:     "mainMenu",
		PrevCommand: "",
		Price:       0,
	}
	commands := []*[]structures.Command{
		{
			{Text: "Магазин", Command: "shop"},
			{Text: "Кабинет", Command: "personalCabinet"}},
		{
			{Text: "Поддержка", Command: "support"},
			{Text: "FAQ", Command: "faq"}},
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
