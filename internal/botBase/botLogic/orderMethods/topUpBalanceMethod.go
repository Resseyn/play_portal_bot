package orderMethods

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/databaseModels"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
)

// TopUpBalance первый метод пополнения баланса, на него ссылаются все методы из магазинов
func TopUpBalance(c telebot.Context) error {

	// =========PARAMS=========
	data := helpingMethods.ParseData(c.Callback().Data)
	fmt.Println(data)
	var user *databaseModels.UserInfo
	data.PrevCommand = ""
	if data.Price != 0 {
		user, _ = databaseModels.Users.GetUser(c.Chat().ID)
		if user.Balance-float64(data.Price) >= 0 {
			commands := [][]structures.Command{{
				{Text: "Вернуться к услуге", Command: structures.Commands["Success"]}}}
			keyboard := helpingMethods.CreateInline(data, commands...)
			c.Send("Вам хватает денег на услугу", keyboard)
			delete(structures.UserStates, c.Chat().ID)
			return nil
		}
	}

	helpingMethods.NewInteraction("awaitingForPrice",
		c.Chat().ID,
		float64(data.Price),
		"",
		[]string{data.Custom})

	picPath := "pkg/utils/data/img/mainMenuImages/Hydra.webp"
	var messageContent string
	var commands [][]structures.Command
	if data.Price == 0 {
		messageContent = fmt.Sprintf("Введите сумму для пополнения от 20₽ и до 20000₽")
		commands = [][]structures.Command{
			{
				{Text: "Вернуться в главное меню", Command: structures.Commands["mainMenu"]}},
		}
	} else {
		data.Price = data.Price - int(user.Balance)
		messageContent = fmt.Sprintf("Вам не хватает %v на балансе\n\nВведите сумму для пополнения от 20₽ и до 20000₽", data.Price)
		commands = [][]structures.Command{
			{
				{Text: fmt.Sprintf("%v₽", data.Price), Command: structures.Commands["createCheck"]}},
			{
				{Text: "Вернуться в главное меню", Command: structures.Commands["mainMenu"]}},
		}
	}
	// =========PARAMS=========
	keyboard := helpingMethods.CreateInline(data, commands...)
	err := c.Edit(&telebot.Photo{
		File:    telebot.FromDisk(picPath),
		Caption: messageContent,
	}, keyboard)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return err
	}
	return nil
}
