package sucessfulPayments

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/databaseModels"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
)

//TODO: СДЕЛАТЬ ОГРОМНЫЙ ХЕНДЛЕР, КОТОРЫЙ БУДЕТ ОБРАБАТЫВАТЬ ВСЕ ЗАПРОСЫ С КЛАВЫ. ДЛЯ КАЖДОГО ХЕНДЛЕРА ("spotifyHandler",..)
//БУДУТ ЗАДАВАТЬСЯ ПАРАМЕТРЫ, ТИПО СКОЛЬКО СООБЩЕНИЙ, КАКОЙ В НИХ ТЕКСТ, НА ОСНОВЕ ДАННЫХ ПАРАМЕТРОВ БУДЕТ ИДТИ ОБРАБОТКА
//ВСЕ ЧЕРЕЗ ОДНО!!!!!

// TODO: переименовать эту хуйню в OrderInfoHandler
// SpotifySuccessPayment - глобальный хендлер для сбора необходимых данный для выполнения заказа юреза, если таких нет - обработка заказа
func SpotifySuccessPayment(c telebot.Context) error {

	delete(structures.UserOrders, c.Chat().ID)
	var params []string

	// =========PARAMS=========
	if c.Callback() != nil { //когда короче первое сообщение там чета
		c.Delete()
		user, _ := databaseModels.Users.GetUser(c.Chat().ID)
		balance := user.Balance
		data := helpingMethods.ParseData(c.Callback().Data)
		params = structures.Parameters[structures.Handlers[data.Custom]]
		if balance-float64(structures.Prices[data.Custom]) >= 0 {
			if len(params) == 0 {
				err := KeySuccess(c)
				if err != nil {
					return err
				}
				return nil
			} else {
				helpingMethods.NewInteraction(
					structures.Handlers[data.Custom],
					c.Chat().ID,
					float64(data.Price),
					data.Custom,
					make([]string, len(params)))
			}
		} else {
			c.Send("Вам не хватает баланса на услугу")
			return nil
		}

	} else {
		params = structures.Parameters[structures.Handlers[structures.UserStates[c.Chat().ID].Order]]
	}
	picPath := "pkg/utils/data/img/shopImages/servicesImages/spotify/spotifySuccess.jpeg"
	messageContent := params[structures.UserStates[c.Chat().ID].Step]

	commands := [][]structures.Command{
		{{Text: "Отмена", Command: structures.Commands["mainMenu"]}}}

	// =========PARAMS=========

	msg := &telebot.Photo{
		File:    telebot.FromDisk(picPath),
		Caption: messageContent,
	}
	keyboard := helpingMethods.CreateInline(&structures.MessageData{}, commands...)
	err := c.Send(msg, &telebot.SendOptions{
		ParseMode:   telebot.ModeHTML,
		ReplyMarkup: keyboard,
	})
	return err
}

func KeySuccess(c telebot.Context) error {
	data := helpingMethods.ParseData(c.Callback().Data)
	key, _ := databaseModels.Keys.GetKey(data.Custom)
	if key == nil || !key.Avaliable {
		c.Send("На данный момент ключей нет, попробуйте позже")
		return nil
	}

	newOrderID := helpingMethods.RandStringRunes(16)
	_, err := databaseModels.Orders.CreateOrder(c.Chat().ID, newOrderID, structures.Prices[data.Custom], data.Custom)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return err
	}
	_, err = databaseModels.Users.ConsumeBalance(c.Chat().ID, structures.Prices[data.Custom])
	if err != nil {
		loggers.ErrorLogger.Println(err)
		c.Bot().Send(telebot.ChatID(c.Chat().ID), "произошла ошибка в бд")
		return err
	}
	_, err = databaseModels.Orders.OrderIsDone(newOrderID)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return err
	}

	c.Send(fmt.Sprintf("Ваш ключ: `%v`", key.Key), telebot.ModeMarkdown)
	//TODO: ну типо гайды условные как вводить ключ за тот или иной товар, тоже + словарь с гойдами
	return nil
}
