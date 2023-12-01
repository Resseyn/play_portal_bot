package sucessfulPayments

import (
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/databaseModels"
	"play_portal_bot/pkg/utils/structures"
)

// SpotifySuccessPayment аквивируется после успешной оплаты юзера, у которого в чеке заказа кастом указан как spotifySuccess
func SpotifySuccessPayment(c telebot.Context) error {

	// =========PARAMS=========
	var data *structures.MessageData
	if c.Callback() != nil {
		data = helpingMethods.ParseData(c.Callback().Data)
	} else {
		data = &structures.MessageData{}
	}
	if _, ok := structures.UserStates[c.Chat().ID]; !ok {
		user, _ := databaseModels.Users.GetUser(c.Chat().ID)
		balance := user.Balance
		if balance-structures.Prices[data.Command] >= 0 {
			structures.UserStates[c.Chat().ID] = &structures.UserInteraction{
				IsInteracting: true,
				Type:          "spotifyHandler",
				Step:          0,
				DataCase:      make([]string, 2),
				Price:         structures.Prices[data.Command],
			}
		} else {
			c.Send("Вам не хватает баланса на услугу")
			return nil
		}
	}

	picPath := "pkg/utils/data/img/shopImages/servicesImages/spotify/spotifySuccess.jpeg"

	var messageContent string
	switch structures.UserStates[c.Chat().ID].Step {
	case 0:
		messageContent = "Введите логин от Spotify"
		c.Delete()
	case 1:
		messageContent = "Введите пароль от Spotify"
	}

	commands := [][]structures.Command{
		{{Text: "Отмена", Command: structures.Commands["mainMenu"]}}}

	data.PrevCommand = "" //TODO: а нужна ли тут дата ваще?Ж???????
	// =========PARAMS=========

	msg := &telebot.Photo{
		File:    telebot.FromDisk(picPath),
		Caption: messageContent,
	}
	keyboard := helpingMethods.CreateInline(data, commands...)
	err := c.Send(msg, &telebot.SendOptions{
		ParseMode:   telebot.ModeHTML,
		ReplyMarkup: keyboard,
	})
	return err
}
