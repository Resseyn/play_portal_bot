package botCommands

import (
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/pkg/utils/structures"
)

// BotStart КОМАНДА СТАРТА, ПО СОВМЕСТИТЕЛЬСТВУ ВЫВОД ГЛАВНОГО МЕНЮ
// func Start(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
//
//		// =========PARAMS=========
//		chatID := update.Message.Chat.ID
//		picPath := "pkg/utils/data/img/pngtree-isolated-cat-on-white-background-png-image_7094927.png"
//		messageContent := "МАГАЗИН ГИДРА"
//		messageData := &structures.MessageData{
//			MessageID:   update.Message.MessageID,
//			ChatID:      update.Message.Chat.ID,
//			Command:     "mainMenu",
//			PrevCommand: "",
//		}
//		fmt.Println(messageData)
//		positions := []int{2, 2}
//		commands := &[]structures.Command{
//			{Text: "Магазин", Command: "showShop"},
//			{Text: "Кабинет", Command: "showPersonalArea"},
//			{Text: "Поддержка", Command: "showSupport"},
//			{Text: "FAQ", Command: "showFAQ"},
//		}
//		// =========PARAMS=========
//
//		msg := helpingMethods.CreateMessage(chatID, picPath, messageContent, commands, messageData, positions)
//		message, err := bot.Send(msg)
//		if err != nil {
//			loggers.ErrorLogger.Println(err.Error())
//		}
//		fmt.Println(message.MessageID, "MESSAGEID SHOP")
//		newKB := helpingMethods.CreateInline(&structures.MessageData{
//			MessageID:   message.MessageID,
//			ChatID:      message.Chat.ID,
//			Command:     "mainMenu",
//			PrevCommand: "",
//		}, []int{2, 2}, *commands...)
//		newKBConf := tgbotapi.NewEditMessageReplyMarkup(chatID, message.MessageID, *newKB)
//		_, err = bot.Send(newKBConf)
//		if err != nil {
//			fmt.Println(messageData)
//			loggers.ErrorLogger.Println(err.Error(), "editKBError")
//		}
//	}

func Start(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/pngtree-isolated-cat-on-white-background-png-image_7094927.png"
	messageContent := "МАГАЗИН ГИДРА"
	messageData := &structures.MessageData{
		MessageID:   c.Message().ID,
		ChatID:      c.Chat().ID,
		Command:     "mainMenu",
		PrevCommand: "",
	}
	commands := []*[]structures.Command{
		{
			{Text: "Магазин", Command: "showShop"},
			{Text: "Кабинет", Command: "showPersonalArea"}},
		{
			{Text: "Поддержка", Command: "showSupport"},
			{Text: "FAQ", Command: "showFAQ"}},
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
