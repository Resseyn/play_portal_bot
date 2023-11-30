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
	messageContent := fmt.Sprintf("%v не понравилось чето гандону", data.Custom)
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

	err := helpingMethods.SendToModers(c, msg, keyboard)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return err
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

	interactionChatID, _ := strconv.Atoi(data.Custom)

	if state, ok := structures.UserStates[int64(interactionChatID)]; ok && state.Type == "moderatorDialog" {
		c.Send("Другой модер пиздит с гандоном")
		return nil
	}

	currentInteraction := &structures.UserInteraction{
		IsInteracting: true,
		Type:          "moderatorDialog",
		DataCase:      []string{strconv.FormatInt(c.Chat().ID, 10)}, //representing moder
	}
	structures.UserStates[int64(interactionChatID)] = currentInteraction

	currentModerInteraction := &structures.UserInteraction{
		IsInteracting: true,
		Type:          "moderatorDialog",
		DataCase:      []string{strconv.FormatInt(int64(interactionChatID), 10)}, //representing user
	}
	structures.UserStates[c.Chat().ID] = currentModerInteraction
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
	_, err = c.Bot().Send(telebot.ChatID(int64(interactionChatID)), msg, keyboard)
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
			Command:     structures.Commands["mainMenu"],
			PrevCommand: "",
			Price:       0,
			Custom:      "",
		}
		messageData2 := &structures.MessageData{
			Command:     structures.Commands["mainMenu"],
			PrevCommand: "",
			Price:       0,
			Custom:      "",
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

	interactionChatID, _ := strconv.Atoi(data.Custom)

	msg := &telebot.Photo{
		File:    telebot.FromDisk(picPath),
		Caption: "Вы прекратили диалог",
	}
	convFrom, _ := strconv.Atoi(structures.UserStates[int64(interactionChatID)].DataCase[0])

	delete(structures.UserStates, c.Chat().ID)
	delete(structures.UserStates, int64(convFrom))

	keyboard := helpingMethods.CreateInline(data, commands...)

	c.Bot().Send(telebot.ChatID(int64(interactionChatID)), msg, keyboard)

	msg.Caption = "С вами прекратили диалог"
	c.Bot().Send(telebot.ChatID(convFrom), msg, keyboard)
	return nil
}
