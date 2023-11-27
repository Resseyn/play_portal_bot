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
			{Text: "Купить", Command: structures.Commands["buy"]}},
	}
	data := helpingMethods.ParseData(c.Callback().Data)
	data.PrevCommand = structures.Commands["shop_services"]
	data.Price = 332
	// =========PARAMS=========
	//currentInteraction := &structures.UserInteraction{
	//	IsInteracting: true,
	//	Type:          "steam_topUpBalance",
	//	Step:          0,
	//	Price:         332,
	//	DataCase:      make([]string, 2),
	//}
	//structures.UserStates[c.Chat().ID] = currentInteraction
	//fmt.Println(*structures.UserStates[c.Chat().ID])
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
