package botLogic

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/databaseModels"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
	"strconv"
)

func Menu(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/mainMenuImages/Hydra.webp"
	messageContent := "МАГАЗИН ИГР 'ЗМЕЙ ГЕРОИНЫЧ'"
	commands := [][]structures.Command{
		{
			{Text: "Магазин", Command: structures.Commands["shop"]},
			{Text: "Кабинет", Command: structures.Commands["personalCabinet"]}},
		{
			{Text: "Поддержка", Command: structures.Commands["support"]},
			{Text: "FAQ", Command: structures.Commands["faq"]}},
	}
	data := &structures.MessageData{
		Command: structures.Commands["mainMenu"],
	}
	// =========PARAMS=========
	if _, ok := structures.UserStates[c.Chat().ID]; ok {
		if structures.UserStates[c.Chat().ID].Type == "moderDialog" {
			c.Send("даун")
			return nil
		}
	}
	delete(structures.UserStates, c.Chat().ID) //TODO: clear userstates func

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
	messageContent := "МАГАЗИН ИГР 'ЗМЕЙ ГЕРОИНЫЧ'"
	data := helpingMethods.ParseData(c.Callback().Data)
	data.PrevCommand = structures.Commands["mainMenu"]
	commands := [][]structures.Command{
		{
			{Text: "Игровые сервисы", Command: structures.Commands["shop_gameServices"]},
			{Text: "Сервисы", Command: structures.Commands["shop_services"]}},
		{
			{Text: "Pepega(насвай не завезли)", Command: structures.Commands[""]},
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
	user, _ := databaseModels.Users.GetUser(c.Chat().ID)
	picPath := "pkg/utils/data/img/mainMenuImages/lcImage.jpeg"
	messageContent := fmt.Sprintf("Общие нары, твое погоняло - %v\n\nБаланс - %v рублей", c.Chat().ID, user.Balance)
	data := helpingMethods.ParseData(c.Callback().Data)
	data.PrevCommand = structures.Commands["mainMenu"]
	commands := [][]structures.Command{
		{
			{Text: "Пополнить баланс💘", Command: structures.Commands["topUpBalance"]}},
		{
			{Text: "Использовать промокод❌", Command: structures.Commands[""]}},
		{
			{Text: "История⚜️", Command: structures.Commands["history"]}},
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
	data.PrevCommand = structures.Commands["mainMenu"]
	commands := [][]structures.Command{
		{
			{Text: "Создать тикет", Command: structures.Commands["createTicket"]}}}
	data.Custom = strconv.Itoa(int(c.Chat().ID))
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
	data.PrevCommand = structures.Commands["mainMenu"]
	commands := [][]structures.Command{
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

func ShowHistory(c telebot.Context) error {
	// =========PARAMS=========
	picPath := "pkg/utils/data/img/mainMenuImages/faq.png"
	messageContent := "Какую историю желаете посмотреть?"
	data := helpingMethods.ParseData(c.Callback().Data)
	data.PrevCommand = structures.Commands["personalCabinet"]
	commands := [][]structures.Command{
		{
			{Text: "Историю пополнений", Command: structures.Commands["historyTOP"]}},
		{
			{Text: "Историю покупок", Command: structures.Commands["historyBUY"]},
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

func ShowHistoryTOP(c telebot.Context) error {
	// =========PARAMS=========
	picPath := "pkg/utils/data/img/mainMenuImages/faq.png"
	messageContent := databaseModels.Orders.ShowOrdersHistory(c.Chat().ID, false)
	data := helpingMethods.ParseData(c.Callback().Data)
	data.PrevCommand = structures.Commands["personalCabinet"]
	commands := [][]structures.Command{
		{{Text: "Историю покупок", Command: structures.Commands["historyBUY"]}}}
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
func ShowHistoryBUY(c telebot.Context) error {
	// =========PARAMS=========
	picPath := "pkg/utils/data/img/mainMenuImages/faq.png"
	messageContent := databaseModels.Orders.ShowOrdersHistory(c.Chat().ID, true)
	data := helpingMethods.ParseData(c.Callback().Data)
	data.PrevCommand = structures.Commands["personalCabinet"]
	commands := [][]structures.Command{
		{
			{Text: "Историю пополнений", Command: structures.Commands["historyTOP"]}},
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
