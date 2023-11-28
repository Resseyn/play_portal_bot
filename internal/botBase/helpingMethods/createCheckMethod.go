package helpingMethods

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
)

func CreateCheck(c telebot.Context) error {

	// =========PARAMS=========
	var data *structures.MessageData
	if c.Callback() != nil {
		data = ParseData(c.Callback().Data)
		err := c.Delete()
		if err != nil {
			loggers.ErrorLogger.Println(err)
			return err
		}
	} else {
		data = &structures.MessageData{
			Command:     structures.UserStates[c.Chat().ID].Type,
			PrevCommand: "",
			Price:       structures.UserStates[c.Chat().ID].Price,
			Custom:      "",
		}
	}

	fmt.Println(data, "НУЖНАЯ ДАТА")
	picPath := "pkg/utils/data/img/mainMenuImages/Hydra.webp"
	messageContent := fmt.Sprintf("Счёт на %v рублей создан, жмите по кнопке ниже, чтобы оплатить удобным вам способом.\n\nДоговор оферты ссылочку ага", data.Price)

	commands := [][]structures.Command{ //TODO:клавиши должны быть с ссылками тут, подумать че делать
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

	if currentState, ok := structures.UserStates[c.Chat().ID]; ok {
		currentState.Type = currentState.DataCase[0]
	}
	fmt.Println(structures.UserStates[c.Chat().ID], "ИЗМЕНЕННАЯ")

	keyboard := CreateInline(data, commands...)
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
