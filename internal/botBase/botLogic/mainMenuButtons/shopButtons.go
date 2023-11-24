package mainMenuButtons

import (
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
)

func GameServices(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/shopImages/gameServices.jpg"
	messageContent := "Выберите категорию"
	commands := []*[]structures.Command{
		{
			{Text: "Steam", Command: "shop_gameServices_steam"},
			{Text: "Xbox/Microsoft", Command: ""}},
		{
			{Text: "Playstation", Command: ""},
			{Text: "ИГРЫ", Command: ""}},
	}
	data := helpingMethods.ParseData(c.Callback().Data)
	data.PrevCommand = "shop"
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
func Services(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/shopImages/gameServices.jpg"
	messageContent := "Выберите категорию"
	commands := []*[]structures.Command{
		{
			{Text: "Подписка Twitch", Command: ""},
			{Text: "Spotify", Command: "spotify"}},
		{
			{Text: "Подписка 7TV", Command: ""},
			{Text: "AppStore", Command: ""}},
		{
			{Text: "Adobe", Command: ""}},
	}
	data := helpingMethods.ParseData(c.Callback().Data)
	data.PrevCommand = "shop"
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
