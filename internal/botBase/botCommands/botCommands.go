package botCommands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
)

// BotStart КОМАНДА СТАРТА, ПО СОВМЕСТИТЕЛЬСТВУ ВЫВОД ГЛАВНОГО МЕНЮ
func BotStart(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	picPath := "pkg/utils/data/img/pngtree-isolated-cat-on-white-background-png-image_7094927.png"
	messageContent := "МАГАЗИН ГИДРА"
	messageData := &structures.MessageData{
		MessageID:   update.Message.MessageID,
		ChatID:      update.Message.Chat.ID,
		Command:     "start",
		PrevCommand: "",
	}
	rows := 2
	columns := 2
	commands := &[]structures.Command{
		{Text: "Магазин", Command: "mainMenu"},
		{Text: "Кабинет", Command: "showPersonalArea"},
		{Text: "Поддержка", Command: "showSupport"},
		{Text: "FAQ", Command: "showFAQ"},
	}
	msg := helpingMethods.CreateMessage(chatID, picPath, messageContent, commands, messageData, rows, columns)
	_, err := bot.Send(msg)
	if err != nil {
		loggers.ErrorLogger.Println(err.Error())
	}
}
