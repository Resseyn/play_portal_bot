package supportMethods

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/botBase/botCommands"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
	"strconv"
)

func CreateTicket(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/shopImages/gameServices.jpg"
	data := helpingMethods.ParseData(c.Callback().Data)
	messageContent := fmt.Sprintf("%v не понравилось чето гандону", data.ChatID)
	commands := [][]structures.Command{
		{
			{Text: "Ответить", Command: structures.Commands["respondToTicket"]}},
	}
	data.PrevCommand = ""
	// =========PARAMS=========

	msg := &telebot.Photo{
		File:    telebot.FromDisk(picPath),
		Caption: messageContent,
	}
	keyboard := helpingMethods.CreateInline(data, commands...)
	for _, moderator := range structures.Moderators {
		moderChat, _ := strconv.Atoi(moderator)
		_, err := c.Bot().Send(telebot.ChatID(moderChat), msg, keyboard)
		if err != nil {
			loggers.ErrorLogger.Println(err)
			return err
		}
	}
	c.Send("Билет создан, ожидайте модератора")
	return nil
}

func RespondToTicket(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/shopImages/gameServices.jpg"
	data := helpingMethods.ParseData(c.Callback().Data)
	commands := [][]structures.Command{
		{
			{Text: "Закончить диалог с хуйланом", Command: structures.Commands["endTicket"]}},
	}
	data.PrevCommand = ""
	// =========PARAMS=========

	if structures.UserStates[data.ChatID].IsInteracting {
		c.Send("Другой модер пиздит с гандоном")
		return nil
	}

	currentInteraction := &structures.UserInteraction{
		IsInteracting: true,
		Type:          "moderatorDialog",
		DataCase:      []string{strconv.FormatInt(c.Chat().ID, 10)}, //representing moder
	}
	structures.UserStates[data.ChatID] = currentInteraction

	currentModerInteraction := &structures.UserInteraction{
		IsInteracting: true,
		Type:          "moderatorDialog",
		DataCase:      []string{strconv.FormatInt(data.ChatID, 10)}, //representing user
	}
	structures.UserStates[c.Chat().ID] = currentModerInteraction
	fmt.Println(currentInteraction)
	fmt.Println(currentModerInteraction)
	msg := &telebot.Photo{
		File:    telebot.FromDisk(picPath),
		Caption: "Вы начали диалог, пропишите /end для окончания или нажмите на кнопку",
	}
	keyboard := helpingMethods.CreateInline(data, commands...)
	_, err := c.Bot().Send(telebot.ChatID(c.Chat().ID), msg, keyboard)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return err
	}
	msg.Caption = "С вами начали диалог, пропишите /end для окончания или нажмите на кнопку"
	_, err = c.Bot().Send(telebot.ChatID(data.ChatID), msg, keyboard)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return err
	}
	return nil
}

func EndTicket(c telebot.Context) error {
	if _, ok := structures.UserStates[c.Chat().ID]; !ok {
		return botCommands.Start(c)
	}
	if c.Callback() == nil {
		convFrom, _ := strconv.Atoi(structures.UserStates[c.Chat().ID].DataCase[0])
		messageData1 := &structures.MessageData{
			MessageID:   c.Message().ID, //бесполезная хуйня
			ChatID:      c.Chat().ID,
			Command:     structures.Commands["mainMenu"],
			PrevCommand: "",
			Price:       0,
		}
		messageData2 := &structures.MessageData{
			MessageID:   c.Message().ID, //бесполезная хуйня
			ChatID:      int64(convFrom),
			Command:     structures.Commands["mainMenu"],
			PrevCommand: "",
			Price:       0,
		}
		picPath := "pkg/utils/data/img/shopImages/gameServices.jpg"
		commands := [][]structures.Command{
			{
				{Text: "Назад в главное меню", Command: structures.Commands["mainMenu"]},
			}}

		delete(structures.UserStates, c.Chat().ID)
		delete(structures.UserStates, int64(convFrom))

		msg := &telebot.Photo{
			File:    telebot.FromDisk(picPath),
			Caption: "Вы прекратили диалог",
		}
		keyboard1 := helpingMethods.CreateInline(messageData1, commands...)
		c.Bot().Send(telebot.ChatID(c.Chat().ID), msg, keyboard1)

		msg.Caption = "С вами прекратили диалог"
		keyboard2 := helpingMethods.CreateInline(messageData2, commands...)
		c.Bot().Send(telebot.ChatID(convFrom), msg, keyboard2)
		return nil
	}
	// =========PARAMS=========
	picPath := "pkg/utils/data/img/shopImages/gameServices.jpg"
	data := helpingMethods.ParseData(c.Callback().Data)
	commands := [][]structures.Command{
		{
			{Text: "Назад в главное меню", Command: structures.Commands["mainMenu"]}},
	}
	data.PrevCommand = ""
	// =========PARAMS=========

	msg := &telebot.Photo{
		File:    telebot.FromDisk(picPath),
		Caption: "Вы прекратили диалог",
	}
	convFrom, _ := strconv.Atoi(structures.UserStates[data.ChatID].DataCase[0])

	delete(structures.UserStates, c.Chat().ID)
	delete(structures.UserStates, int64(convFrom))

	keyboard := helpingMethods.CreateInline(data, commands...)

	c.Bot().Send(telebot.ChatID(data.ChatID), msg, keyboard)

	msg.Caption = "С вами прекратили диалог"
	c.Bot().Send(telebot.ChatID(convFrom), msg, keyboard)
	return nil
}
