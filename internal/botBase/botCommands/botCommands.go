package botCommands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
)

// BotStart КОМАНДА СТАРТА, ПО СОВМЕСТИТЕЛЬСТВУ ВЫВОД ГЛАВНОГО МЕНЮ
func BotStart(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {

	// =========PARAMS=========
	chatID := update.Message.Chat.ID
	picPath := "pkg/utils/data/img/pngtree-isolated-cat-on-white-background-png-image_7094927.png"
	messageContent := "МАГАЗИН ГИДРА"
	messageData := &structures.MessageData{
		MessageID:   update.Message.MessageID,
		ChatID:      update.Message.Chat.ID,
		Command:     "start",
		PrevCommand: "",
	}
	fmt.Println(messageData)
	positions := []int{2, 2}
	commands := &[]structures.Command{
		{Text: "Магазин", Command: "mainMenu"},
		{Text: "Кабинет", Command: "showPersonalArea"},
		{Text: "Поддержка", Command: "showSupport"},
		{Text: "FAQ", Command: "showFAQ"},
	}
	// =========PARAMS=========

	msg := helpingMethods.CreateMessage(chatID, picPath, messageContent, commands, messageData, positions)
	message, err := bot.Send(msg)
	if err != nil {
		loggers.ErrorLogger.Println(err.Error())
	}
	fmt.Println(message.MessageID, "MESSAGEID SHOP")
	newKB := helpingMethods.CreateInline(&structures.MessageData{
		MessageID:   message.MessageID,
		ChatID:      update.Message.Chat.ID,
		Command:     "start",
		PrevCommand: "",
	}, []int{4}, *commands...)
	newKBConf := tgbotapi.NewEditMessageReplyMarkup(chatID, message.MessageID, *newKB)
	_, err = bot.Send(newKBConf)
	if err != nil {
		loggers.ErrorLogger.Println(err.Error(), "editKBError")
	}
}
