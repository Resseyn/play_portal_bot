package sucessfulPayments

import (
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/databaseModels"
	"play_portal_bot/pkg/utils/structures"
)

//TODO: СДЕЛАТЬ ОГРОМНЫЙ ХЕНДЛЕР, КОТОРЫЙ БУДЕТ ОБРАБАТЫВАТЬ ВСЕ ЗАПРОСЫ С КЛАВЫ. ДЛЯ КАЖДОГО ХЕНДЛЕРА ("spotifyHandler",..)
//БУДУТ ЗАДАВАТЬСЯ ПАРАМЕТРЫ, ТИПО СКОЛЬКО СООБЩЕНИЙ, КАКОЙ В НИХ ТЕКСТ, НА ОСНОВЕ ДАННЫХ ПАРАМЕТРОВ БУДЕТ ИДТИ ОБРАБОТКА
//ВСЕ ЧЕРЕЗ ОДНО!!!!!

// SpotifySuccessPayment аквивируется после успешной оплаты юзера, у которого в чеке заказа кастом указан как spotifySuccess
func SpotifySuccessPayment(c telebot.Context) error {
	delete(structures.UserRedirectsAndOrders, c.Chat().ID)

	//Params[Handler[data.Custom]] = []string{"Введите логин от спот", "введите пароль"} //TODO: два словаря новых ага
	// =========PARAMS=========
	if c.Callback() != nil { //когда короче первое сообщение там чета
		data := helpingMethods.ParseData(c.Callback().Data)
		user, _ := databaseModels.Users.GetUser(c.Chat().ID)
		balance := user.Balance
		if balance-float64(data.Price) >= 0 {
			helpingMethods.NewInteraction(
				"spotifyHandler",
				c.Chat().ID,
				float64(data.Price),
				data.Custom,
				make([]string, 2)) //make ([] string, len(params))
		} else {
			c.Send("Вам не хватает баланса на услугу")
			return nil
		}

	}
	picPath := "pkg/utils/data/img/shopImages/servicesImages/spotify/spotifySuccess.jpeg"
	//for i, text := range params{if i == step && i - 1 != len(params) {messageContent = text; c.Delete } else..... }
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

	// =========PARAMS=========

	msg := &telebot.Photo{
		File:    telebot.FromDisk(picPath),
		Caption: messageContent,
	}
	keyboard := helpingMethods.CreateInline(&structures.MessageData{}, commands...)
	err := c.Send(msg, &telebot.SendOptions{
		ParseMode:   telebot.ModeHTML,
		ReplyMarkup: keyboard,
	})
	return err
}

func KeySuccess(c telebot.Context) error {
	//TODO:
	return nil
}
