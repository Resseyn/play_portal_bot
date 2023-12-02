package orderMethods

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
)

// CreateCheck второй метод пополнения, который уже создает чек, когда пользователь определился с суммой пополнения,
// от него идут все платежки
func CreateCheck(c telebot.Context) error {
	if currentState, ok := structures.UserStates[c.Chat().ID]; ok {
		if len(currentState.DataCase) != 0 {
			currentState.Type = currentState.DataCase[0]
		}
	} else {
		c.Send("Произошла ошибка")
		return nil
	}

	// =========PARAMS=========
	var data *structures.MessageData
	if c.Callback() != nil {
		data = helpingMethods.ParseData(c.Callback().Data)
		err := c.Delete()
		if err != nil {
			loggers.ErrorLogger.Println(err)
			return err
		}
	} else {
		data = &structures.MessageData{
			Command:     structures.UserStates[c.Chat().ID].Type,
			PrevCommand: "",
			Price:       int(structures.UserStates[c.Chat().ID].Price),
			Custom:      structures.UserStates[c.Chat().ID].Type,
		}
	}

	picPath := "pkg/utils/data/img/mainMenuImages/Hydra.webp"
	messageContent := fmt.Sprintf("Счёт на %v рублей создан, жмите по кнопке ниже, чтобы оплатить удобным вам способом.\n\nДоговор оферты ссылочку ага", data.Price)

	commands := [][]structures.Command{
		{
			{Text: "PayPalych", Command: structures.Commands["createPayPalychBill"]}},
		{
			{Text: fmt.Sprintf("PayPalych(%v$)", data.Price), Command: ""},
			{Text: "LavaRu", Command: ""}},
		{
			{Text: "Вернуться в главное меню", Command: structures.Commands["mainMenu"]}},
	}
	data.PrevCommand = ""
	// =========PARAMS=========

	delete(structures.UserStates, c.Chat().ID)
	keyboard := helpingMethods.CreateInline(data, commands...)
	err := c.Send(&telebot.Photo{
		File:    telebot.FromDisk(picPath),
		Caption: messageContent,
	}, keyboard)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return err
	}
	return nil
}
