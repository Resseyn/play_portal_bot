package servicesButtons

import (
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
)

func Spotify(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/shopImages/gameServices.jpg"
	messageContent := "• Аккаунты: купить новый аккаунт Steam.\n\n• Пополнить баланс: автоматическая система пополнения Steam баланса для России, Казахстана, Украины.\n\n жирный шрифт добавить"
	commands := []*[]structures.Command{
		{
			{Text: "Пополнить баланс", Command: "shop_gameServices_steam_topUpBalance"}},
		{
			{Text: "Аккаунты", Command: "steam_accounts"}},
	}
	data := helpingMethods.ParseData(c.Callback().Data)
	data.PrevCommand = "gameServices"
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
