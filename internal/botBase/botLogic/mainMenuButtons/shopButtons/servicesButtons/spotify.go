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
			{Text: "Spotify Individual 3 месяца", Command: structures.Commands["spotify_individual_3"]}},
		{
			{Text: "Spotify Individual 6 месяцев", Command: structures.Commands["spotify_individual_6"]}},
		{
			{Text: "Spotify Individual 12 месяцев", Command: structures.Commands["spotify_individual_12"]}},
		{
			{Text: "Spotify DUO 1 месяц", Command: structures.Commands["spotify_duo_1"]}},
		{
			{Text: "Spotify Family 1 месяц", Command: structures.Commands["spotify_family_1"]}},
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
	data.Custom = "spoa"
	data.PrevCommand = structures.Commands["shop_services"]
	data.Price = int(structures.Prices[data.Custom])
	// =========PARAMS=========

	//TODO: глобальный рефрактор, спотифайсаксес теперть просто саксес Инпутпэймент
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
func SpotifyIndividual3(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/shopImages/servicesImages/spotify/spotify_individual_3.jpg"
	messageContent := "Выберите товар:"
	commands := [][]structures.Command{
		{
			{Text: "Купить", Command: structures.Commands["topUpBalance"]}},
	}
	data := helpingMethods.ParseData(c.Callback().Data)
	data.Custom = "spob"
	data.PrevCommand = structures.Commands["shop_services"]
	data.Price = int(structures.Prices[data.Custom])
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
func SpotifyIndividual6(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/shopImages/servicesImages/spotify/spotify_individual_6.jpg"
	messageContent := "Выберите товар:"
	commands := [][]structures.Command{
		{
			{Text: "Купить", Command: structures.Commands["topUpBalance"]}},
	}
	data := helpingMethods.ParseData(c.Callback().Data)
	data.Custom = "spoc"
	data.PrevCommand = structures.Commands["shop_services"]
	data.Price = int(structures.Prices[data.Custom])
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
func SpotifyIndividual12(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/shopImages/servicesImages/spotify/spotify_individual_12.jpg"
	messageContent := "Выберите товар:"
	commands := [][]structures.Command{
		{
			{Text: "Купить", Command: structures.Commands["topUpBalance"]}},
	}
	data := helpingMethods.ParseData(c.Callback().Data)
	data.Custom = "spod"
	data.PrevCommand = structures.Commands["shop_services"]
	data.Price = int(structures.Prices[data.Custom])
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
func SpotifyDUO1(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/shopImages/servicesImages/spotify/spotify_DUO_1.jpg"
	messageContent := "Выберите товар:"
	commands := [][]structures.Command{
		{
			{Text: "Купить", Command: structures.Commands["topUpBalance"]}},
	}
	data := helpingMethods.ParseData(c.Callback().Data)
	data.Custom = "spoe"
	data.PrevCommand = structures.Commands["shop_services"]
	data.Price = int(structures.Prices[data.Custom])
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
func SpotifyFAM1(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/shopImages/servicesImages/spotify/spotify_family_1.jpg"
	messageContent := "Выберите товар:"
	commands := [][]structures.Command{
		{
			{Text: "Купить", Command: structures.Commands["topUpBalance"]}},
	}
	data := helpingMethods.ParseData(c.Callback().Data)
	data.Custom = "spof"
	data.PrevCommand = structures.Commands["shop_services"]
	data.Price = int(structures.Prices[data.Custom])
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
