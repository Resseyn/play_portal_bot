package helpingMethods

//
//import (
//	"gopkg.in/telebot.v3"
//	"play_portal_bot/internal/botBase/keys"
//	"play_portal_bot/internal/loggers"
//)
//
//func CreateCheck(c telebot.Context) error {
//
//	// =========PARAMS=========
//	//dollarCourse := 90
//
//	data := ParseData(c.Callback().Data)
//	picPath := "pkg/utils/data/img/mainMenuImages/Hydra.webp"
//	messageContent := "Счёт на 332 рублей создан, жмите по кнопке ниже, чтобы оплатить удобным вам способом.\n\nДоговор оферты ссылочку ага"
//	//commands := []*[]structures.Command{ //TODO:клавиши должны быть с ссылками тут, подумать че делать
//	//	{
//	//		{Text: "PayPalych", Command: ""}},
//	//	{
//	//		{Text: fmt.Sprintf("PayPalych(%v$)", data.Price/dollarCourse), Command: ""},
//	//		{Text: "LavaRu", Command: ""}},
//	//	{
//	//		{Text: "Изменить сумму", Command: data.PrevCommand}},
//	//	{
//	//		{Text: "Вернуться в главное меню", Command: "mainMenu"}},
//	//}
//	data.PrevCommand = ""
//	// =========PARAMS=========
//
//	inv := telebot.Invoice{Photo: &telebot.Photo{
//		File:    telebot.FromDisk(picPath),
//		Caption: messageContent,
//	}, Title: "test", Description: "testDESC", Payload: "sber", Token: keys.SberPaymentToken, Currency: "RUB", Prices: []telebot.Price{telebot.Price{
//		Label: "zaplatit", Amount: 22800,
//	},
//	}, Start: "sdad"}
//	//keyboard := CreateInline(data, commands...)
//	//err := c.Edit(&telebot.Photo{
//	//	File:    telebot.FromDisk(picPath),
//	//	Caption: messageContent,
//	//}, keyboard)
//	//err := c.Edit(inv)
//	_, err := inv.Send(c.Bot(), telebot.ChatID(c.Chat().ID), nil)
//	if err != nil {
//		loggers.ErrorLogger.Println(err)
//		return err
//	}
//	return nil
//}
