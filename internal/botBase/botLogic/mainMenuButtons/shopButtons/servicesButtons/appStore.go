package servicesButtons

import (
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
)

func AppStore(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/shopImages/servicesImages/appStore.jpg"
	messageContent := "Выберите товар:"
	commands := [][]structures.Command{
		{
			{Text: "Ключ AppStore 500 руб", Command: structures.Commands["appStore500"]}},
		{
			{Text: "Ключ AppStore 1000 руб", Command: structures.Commands["appStore1000"]}},
		{
			{Text: "Ключ AppStore 1500 руб", Command: structures.Commands["appStore1500"]}},
		{
			{Text: "Ключ AppStore 3000 руб", Command: structures.Commands["appStore3000"]}},
		{
			{Text: "Ключ AppStore 9000 руб", Command: structures.Commands["appStore9000"]}},
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
func AppStore500key(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/shopImages/servicesImages/appStore/appStore500.jpg"
	messageContent := "<b>Товар:</b> Ключ App Store 500 руб\n<b>Цена:</b> 689₽\n\n<b>Описание:</b> Подарочная карта оплаты AppStore/iTunes. \nС помощью баланса можно оплачивать подписки/покупать товары в играх и приложениях.\n\n⚠️Ключ можно активировать только на учетной записи РФ региона."
	commands := [][]structures.Command{
		{
			{Text: "Купить", Command: structures.Commands["topUpBalance"]}},
	}
	data := helpingMethods.ParseData(c.Callback().Data)
	data.Custom = "app1"
	data.PrevCommand = structures.Commands["shop_services"]
	data.Price = int(structures.Prices[data.Custom])
	// =========PARAMS=========

	structures.UserRedirectsAndOrders[c.Chat().ID] = []string{structures.Commands["keySuccess"], data.Custom}
	keyboard := helpingMethods.CreateInline(data, commands...)
	err := c.Edit(&telebot.Photo{
		File:    telebot.FromDisk(picPath),
		Caption: messageContent,
	}, keyboard, telebot.ParseMode("HTML")) //TODO: проверить как делать жирный шрифт нормально
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return err
	}
	return nil
}
func AppStore1000key(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/shopImages/servicesImages/appStore/appStore1000.jpg"
	messageContent := "Выберите товар:"
	commands := [][]structures.Command{
		{
			{Text: "Купить", Command: structures.Commands["topUpBalance"]}},
	}
	data := helpingMethods.ParseData(c.Callback().Data)
	data.Custom = "app2"
	data.PrevCommand = structures.Commands["shop_services"]
	data.Price = int(structures.Prices[data.Custom])
	// =========PARAMS=========

	structures.UserRedirectsAndOrders[c.Chat().ID] = []string{structures.Commands["keySuccess"], data.Custom}
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
func AppStore1500key(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/shopImages/servicesImages/appStore/appStore1500.jpg"
	messageContent := "Выберите товар:"
	commands := [][]structures.Command{
		{
			{Text: "Купить", Command: structures.Commands["topUpBalance"]}},
	}
	data := helpingMethods.ParseData(c.Callback().Data)
	data.Custom = "app3"
	data.PrevCommand = structures.Commands["shop_services"]
	data.Price = int(structures.Prices[data.Custom])
	// =========PARAMS=========

	structures.UserRedirectsAndOrders[c.Chat().ID] = []string{structures.Commands["keySuccess"], data.Custom}
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
func AppStore3000key(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/shopImages/servicesImages/appStore/appStore3000.jpg"
	messageContent := "Выберите товар:"
	commands := [][]structures.Command{
		{
			{Text: "Купить", Command: structures.Commands["topUpBalance"]}},
	}
	data := helpingMethods.ParseData(c.Callback().Data)
	data.Custom = "app4"
	data.PrevCommand = structures.Commands["shop_services"]
	data.Price = int(structures.Prices[data.Custom])
	// =========PARAMS=========

	structures.UserRedirectsAndOrders[c.Chat().ID] = []string{structures.Commands["keySuccess"], data.Custom}
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
func AppStore9000key(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/shopImages/servicesImages/appStore/appStore9000.jpg"
	messageContent := "Выберите товар:"
	commands := [][]structures.Command{
		{
			{Text: "Купить", Command: structures.Commands["topUpBalance"]}},
	}
	data := helpingMethods.ParseData(c.Callback().Data)
	data.Custom = "app5"
	data.PrevCommand = structures.Commands["shop_services"]
	data.Price = int(structures.Prices[data.Custom])
	// =========PARAMS=========

	structures.UserRedirectsAndOrders[c.Chat().ID] = []string{structures.Commands["keySuccess"], data.Custom}
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
