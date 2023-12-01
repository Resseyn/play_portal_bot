package helpingMethods

import (
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/databaseModels"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
	"strconv"
)

// CreateOrder создает обработку заказа после подтверждения оплаты и снятия средств короче шоб все ок было
func CreateOrder(c telebot.Context) error {
	if !CheckIfIsInteracting(c.Chat().ID) {
		return nil
	}

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/shopImages/gameServices.jpg"
	data := &structures.MessageData{
		Custom: strconv.Itoa(int(c.Chat().ID)),
		Price:  int(structures.UserStates[c.Chat().ID].Price),
	}
	messageContent := "Заказ CO4DOM81K3 принят в обработку!\n\n📝 Ваш заказ выполняется вручную, возможно оператору понадобится с вами связаться. Поэтому пожалуйста не отключайте уведомления в чате.\n\nНажмите на кнопку ниже, чтобы дать понять оператору, что находитесь в сети и ожидаете заказ! Выполнение заказа не всегда начинается мгновенно из-за возможных очередей, поэтому придётся немного подождать.\n\nСпасибо за терпение ❤️"
	commands := [][]structures.Command{
		{
			{Text: "Ответить", Command: structures.Commands["respondToOrder"]}},
	}
	data.PrevCommand = ""
	// =========PARAMS=========

	msg := &telebot.Photo{
		File:    telebot.FromDisk(picPath),
		Caption: messageContent,
	}
	keyboard := CreateInline(data, commands...)
	c.Send(msg)
	msg.Caption = data.Custom + structures.UserStates[c.Chat().ID].Type + structures.UserStates[c.Chat().ID].DataCase[0] + structures.UserStates[c.Chat().ID].DataCase[1]
	//TODO: убирать юзерстейты блять там где они не нужны, использовать бд в будущем, юзерстейт онли для долгого взаимодействия
	for _, moderator := range structures.Moderators {
		moderChat, _ := strconv.Atoi(moderator)
		_, err := c.Bot().Send(telebot.ChatID(moderChat), msg, keyboard)
		if err != nil {
			loggers.ErrorLogger.Println(err)
			return err
		}
	}
	return nil
}

// RespondToOrder функция для вступления модера в диалог с пользователем
func RespondToOrder(c telebot.Context) error {
	if !CheckIfIsInteracting(c.Chat().ID) {
		return nil
	}

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/shopImages/gameServices.jpg"
	data := ParseData(c.Callback().Data)
	commands := [][]structures.Command{
		{
			{Text: "Закончить заказ с хуйланом", Command: structures.Commands["endOrder"]}},
	}
	data.PrevCommand = ""
	// =========PARAMS=========

	interactionChatID, _ := strconv.Atoi(data.Custom)

	if state, ok := structures.UserStates[int64(interactionChatID)]; ok && state.Type == "moderatorDialog" {
		c.Send("Другой модер занят гандоном")
		return nil
	}

	currentInteraction := &structures.UserInteraction{
		IsInteracting: true,
		Type:          "moderatorDialog",
		DataCase:      []string{strconv.FormatInt(c.Chat().ID, 10)}, //representing user
	}
	structures.UserStates[int64(interactionChatID)] = currentInteraction

	currentModerInteraction := &structures.UserInteraction{
		IsInteracting: true,
		Type:          "moderatorDialog",
		DataCase:      []string{strconv.FormatInt(int64(interactionChatID), 10)}, //representing moder
		Price:         float64(data.Price),
	}
	structures.UserStates[c.Chat().ID] = currentModerInteraction
	msg := &telebot.Photo{
		File:    telebot.FromDisk(picPath),
		Caption: "Вы начали обрботку заказа, пропишите /endOrder для окончания или нажмите на кнопку",
	}
	keyboard := CreateInline(data, commands...)
	_, err := c.Bot().Send(telebot.ChatID(c.Chat().ID), msg, keyboard)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return err
	}
	msg.Caption = "Оператор подключился"
	_, err = c.Bot().Send(telebot.ChatID(int64(interactionChatID)), msg)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return err
	}
	return nil
}

// EndOrder модер выполнил работу и нажимает эту кнопку
func EndOrder(c telebot.Context) error { //TODO: заказ должен будет удаляться из бд, деньги списываться с юзера и т.д
	if !CheckIfIsInteracting(c.Chat().ID) {
		return nil
	}
	convFrom, _ := strconv.Atoi(structures.UserStates[c.Chat().ID].DataCase[0])

	_, err := databaseModels.Users.ConsumeBalance(c.Chat().ID, float64(structures.UserStates[c.Chat().ID].Price))
	if err != nil {
		loggers.ErrorLogger.Println(err)
		c.Bot().Send(telebot.ChatID(convFrom), "произошла ошибка в бд")
		return err
	}
	if c.Callback() == nil {
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
		keyboard1 := CreateInline(messageData1, commands...)
		c.Bot().Send(telebot.ChatID(c.Chat().ID), msg, keyboard1)

		msg.Caption = "С вами прекратили диалог"
		keyboard2 := CreateInline(messageData2, commands...)
		c.Bot().Send(telebot.ChatID(convFrom), msg, keyboard2)
		return nil
	}
	// =========PARAMS=========
	picPath := "pkg/utils/data/img/shopImages/gameServices.jpg"
	data := ParseData(c.Callback().Data)
	commands := [][]structures.Command{
		{
			{Text: "Назад в главное меню", Command: structures.Commands["mainMenu"]}},
	}
	data.PrevCommand = ""
	// =========PARAMS=========

	//interactionChatID, _ := strconv.Atoi(data.Custom)

	msg := &telebot.Photo{
		File:    telebot.FromDisk(picPath),
		Caption: "Вы прекратили диалог",
	}
	//convFrom, _ := strconv.Atoi(structures.UserStates[int64(interactionChatID)].DataCase[0])

	delete(structures.UserStates, c.Chat().ID)
	delete(structures.UserStates, int64(convFrom))

	keyboard := CreateInline(data, commands...)

	c.Bot().Send(telebot.ChatID(int64(convFrom)), msg, keyboard)

	msg.Caption = "Ваш заказ выполнен!!!!!"
	c.Bot().Send(telebot.ChatID(convFrom), msg, keyboard)
	return nil
}
