package servicesButtons

import (
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
)

func Spotify(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/shopImages/servicesImages/spotify.jpg"
	messageContent := "Выберите товар:"
	commands := [][]structures.Command{
		{
			{Text: "Spotify Individual 1 месяц", Command: structures.Commands["spotify_individual_1"]}},
		{
			{Text: "Spotify Individual 3 месяца", Command: structures.Commands["steam_accounts"]}},
		{
			{Text: "Spotify Individual 6 месяцев", Command: structures.Commands["steam_accounts"]}},
		{
			{Text: "Spotify Individual 12 месяцев", Command: structures.Commands["steam_accounts"]}},
		{
			{Text: "Spotify DUO 1 месяц", Command: structures.Commands["steam_accounts"]}},
		{
			{Text: "Spotify Family 1 месяц", Command: structures.Commands["steam_accounts"]}},
	}
	data := helpingMethods.ParseData(c.Callback().Data)
	data.PrevCommand = structures.Commands["shop_services"]
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
func SpotifyIndividual1(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/shopImages/servicesImages/spotify/spotify_individual_1.jpg"
	messageContent := "Выберите товар:"
	commands := [][]structures.Command{
		{
			{Text: "Купить", Command: structures.Commands["topUpBalance"]}},
	}
	data := helpingMethods.ParseData(c.Callback().Data)
	data.Command = structures.Commands["spotifySuccess"]
	data.PrevCommand = structures.Commands["shop_services"]
	data.Price = 332
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

// SpotifySuccessPayment аквивируется после успешной оплаты юзера, у которого в чеке заказа кастом указан как spotifySuccess
func SpotifySuccessPayment(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/shopImages/servicesImages/spotify/spotifySuccess.jpeg"
	var messageContent string
	commands := [][]structures.Command{
		{{Text: "Отмена", Command: structures.Commands["mainMenu"]}}}
	var data *structures.MessageData
	if c.Callback() != nil {
		data = helpingMethods.ParseData(c.Callback().Data)
	} else {
		data = &structures.MessageData{}
	}
	data.PrevCommand = "" //TODO: а нужна ли тут дата ваще?Ж???????
	// =========PARAMS=========

	switch structures.UserStates[c.Chat().ID].Step {
	case 0:
		messageContent = "Введите логин от Spotify"
		c.Delete()
	case 1:
		messageContent = "Введите пароль от Spotify"
	}

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
