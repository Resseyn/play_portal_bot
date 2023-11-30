package helpingMethods

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
)

func TopUpBalance(c telebot.Context) error {

	// =========PARAMS=========
	data := ParseData(c.Callback().Data)
	data.Custom = data.PrevCommand
	NewInteraction("awaitingForPrice", c.Chat().ID, data.Price, []string{data.Custom})
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
		messageContent = fmt.Sprintf("Вам не хватает %v на балансе\n\nВведите сумму для пополнения от 20₽ и до 20000₽", data.Price)
		commands = [][]structures.Command{
			{
				{Text: fmt.Sprintf("%v₽", data.Price), Command: structures.Commands["createCheck"]}},
			{
				{Text: "Вернуться в главное меню", Command: structures.Commands["mainMenu"]}},
		}
	}
	data.PrevCommand = ""
	// =========PARAMS=========
	//TODO:прописать метод пополнения по тексту
	keyboard := CreateInline(data, commands...)
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
func UpdateTopUpBalance(c telebot.Context) error {
	return nil
}
