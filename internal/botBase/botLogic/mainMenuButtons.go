package botLogic

import (
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
)

func Menu(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/mainMenuImages/Hydra.webp"
	messageContent := "МАГАЗИН ИГР 'ГИДРА'"
	commands := []*[]structures.Command{
		{
			{Text: "Магазин", Command: "shop"},
			{Text: "Кабинет", Command: "personalCabinet"}},
		{
			{Text: "Поддержка", Command: "support"},
			{Text: "FAQ", Command: "faq"}},
	}
	data := helpingMethods.ParseData(c.Callback().Data)
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

func Shop(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/mainMenuImages/gettyimages-1067956982.jpg"
	messageContent := "МАГАЗИН ИГР 'ГИДРА'"
	data := helpingMethods.ParseData(c.Callback().Data)
	data.PrevCommand = "mainMenu"
	commands := []*[]structures.Command{
		{
			{Text: "Игровые сервисы", Command: "gameServices"},
			{Text: "Сервисы", Command: "services"}},
		{
			{Text: "Pepega(насвай не завезли)", Command: ""},
		}}
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
func PersonalCabinet(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/mainMenuImages/lcImage.jpeg"
	messageContent := "Общие нары"
	data := helpingMethods.ParseData(c.Callback().Data)
	data.PrevCommand = "mainMenu"
	commands := []*[]structures.Command{
		{
			{Text: "Пополнить баланс💘", Command: ""}},
		{
			{Text: "Использовать промокод❌", Command: ""}},
		{
			{Text: "История покупок⚜️", Command: ""}},
	}
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
func Support(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/mainMenuImages/best-hard-support-dota-2-heroes-1-e1687346780280.jpg"
	messageContent := "Вы можете задать свой вопрос в поддержку создав тикет, но перед этим рекомендуем ознакомиться с нашим FAQ"
	data := helpingMethods.ParseData(c.Callback().Data)
	data.PrevCommand = "mainMenu"
	commands := []*[]structures.Command{
		{
			{Text: "Создать тикет", Command: ""}}}
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
func FAQ(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/mainMenuImages/faq.png"
	messageContent := "Здесь можно почитать ответы на Часто задаваемые вопросы. НУ И ТИПО ССЫЛОЧКУ СЮДА АГА" //сюда ссылку
	data := helpingMethods.ParseData(c.Callback().Data)
	data.PrevCommand = "mainMenu"
	commands := []*[]structures.Command{
		{}}
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