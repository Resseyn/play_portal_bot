package sucessfulPayments

import (
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/databaseModels"
	"play_portal_bot/pkg/utils/structures"
)

//TODO: СДЕЛАТЬ ОГРОМНЫЙ ХЕНДЛЕР, КОТОРЫЙ БУДЕТ ОБРАБАТЫВАТЬ ВСЕ ЗАПРОСЫ С КЛАВЫ. ДЛЯ КАЖДОГО ХЕНДЛЕРА ("spotifyHandler",..)
//БУДУТ ЗАДАВАТЬСЯ ПАРАМЕТРЫ, ТИПО СКОЛЬКО СООБЩЕНИЙ, КАКОЙ В НИХ ТЕКСТ, НА ОСНОВЕ ДАННЫХ ПАРАМЕТРОВ БУДЕТ ИДТИ ОБРАБОТКА
//ВСЕ ЧЕРЕЗ ОДНО!!!!!

// TODO: переименовать эту хуйню в OrderInfoHandler
// SpotifySuccessPayment аквивируется после успешной оплаты юзера, у которого в чеке заказа кастом указан как spotifySuccess
func SpotifySuccessPayment(c telebot.Context) error {

	delete(structures.UserOrders, c.Chat().ID)
	var params []string

	// =========PARAMS=========
	if c.Callback() != nil { //когда короче первое сообщение там чета
		user, _ := databaseModels.Users.GetUser(c.Chat().ID)
		balance := user.Balance
		data := helpingMethods.ParseData(c.Callback().Data)
		params = structures.Parameters[structures.Handlers[data.Custom]]
		if balance-float64(structures.Prices[data.Custom]) >= 0 {
			if len(params) == 0 {
				//keySuccess
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
	//TODO:
	return nil
}
