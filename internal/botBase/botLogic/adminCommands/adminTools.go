package adminCommands

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"math/rand"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/pkg/utils/structures"
)

// мне было супервпадлу делать чтобы не нужно было создавать две отдельных функции
// но оказалось мне похуй и пусть все будет вот так, создание сообщения в botCommands.go
func ShowAdminPanel(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/adminImages/ramzes.jpg"
	messageContent := fmt.Sprintf("админ менюшка (тест %.2f)", rand.Float32()*1.3)
	messageData := &structures.MessageData{
		Command:     structures.Commands["adminPanel"],
		PrevCommand: "",
		Price:       0,
		Custom:      "",
	}
	commands := [][]structures.Command{
		{
			{Text: "Показать репорты", Command: structures.Commands["showReports"]},
			{Text: "Падрачить хуй", Command: structures.Commands["showAdminPanel"]}}}

	// =========PARAMS=========

	keyboard := helpingMethods.CreateInline(messageData, commands...)
	err := c.Edit(&telebot.Photo{
		File:    telebot.FromDisk(picPath),
		Caption: messageContent,
	}, keyboard)

	return err

}

// условно из FAQ идет то самое создание тикета и вот допустим они сюда идут ну я думаю потом придумаем куда это чисто временная прикольная хуйня
func ShowReports(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/adminImages/astronaut.jpg"
	messageContent := "долбаебы понапишут вопросов через факью и они все сюда"
	messageData := &structures.MessageData{
		Command:     structures.Commands["showReports"],
		PrevCommand: structures.Commands["showAdminPanel"],
		Price:       0,
		Custom:      "",
	}
	commands := [][]structures.Command{{}}
	// =========PARAMS=========

	keyboard := helpingMethods.CreateInline(messageData, commands...)
	err := c.Edit(&telebot.Photo{
		File:    telebot.FromDisk(picPath),
		Caption: messageContent,
	}, keyboard)
	return err
}
