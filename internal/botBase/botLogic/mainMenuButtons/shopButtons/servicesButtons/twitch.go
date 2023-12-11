package servicesButtons

//import (
//	"gopkg.in/telebot.v3"
//	"play_portal_bot/internal/botBase/helpingMethods"
//	"play_portal_bot/internal/loggers"
//	"play_portal_bot/pkg/utils/structures"
//)
//
//func Twitch(c telebot.Context) error {
//
//	// =========PARAMS=========
//	picPath := "pkg/utils/data/img/shopImages/servicesImages/twitch.jpg"
//	messageContent := "Выберите товар:"
//	commands := [][]structures.Command{
//		{
//			{Text: "Spotify Individual 1 месяц", Command: structures.Commands["spotify_individual_1"]}},
//		{
//			{Text: "Spotify Individual 3 месяца", Command: structures.Commands["spotify_individual_3"]}},
//		{
//			{Text: "Spotify Individual 6 месяцев", Command: structures.Commands["spotify_individual_6"]}},
//		{
//			{Text: "Spotify Individual 12 месяцев", Command: structures.Commands["spotify_individual_12"]}},
//		{
//			{Text: "Spotify DUO 1 месяц", Command: structures.Commands["spotify_duo_1"]}},
//		{
//			{Text: "Spotify Family 1 месяц", Command: structures.Commands["spotify_family_1"]}},
//	}
//	data := helpingMethods.ParseData(c.Callback().Data)
//	data.PrevCommand = structures.Commands["shop_services"]
//	// =========PARAMS=========
//
//	keyboard := helpingMethods.CreateInline(data, commands...)
//	err := c.Edit(&telebot.Photo{
//		File:    telebot.FromDisk(picPath),
//		Caption: messageContent,
//	}, keyboard)
//	if err != nil {
//		loggers.ErrorLogger.Println(err)
//		return err
//	}
//	return nil
//}
