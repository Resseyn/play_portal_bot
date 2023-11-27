package steamButtons

import (
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
)

func SteamTopUpBalance(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/shopImages/gameServices.jpg"
	messageContent := "Укажите логин Steam аккаунта, на который хотите пополнить баланс: \n \nАвтоматическое пополнение баланса в Steam на аккаунты России, Украины, Казахстана. \n \n⚠️ Обратите внимание! Логин - это то, что вы указываете при входе в Steam. Указав неверные данные, средства уйдут другому пользователю. посмотреть свой логин\n\nПрочитай если уровень STEAM LVL 0"
	commands := [][]structures.Command{{}}
	data := helpingMethods.ParseData(c.Callback().Data)
	data.PrevCommand = structures.Commands["steam"]
	// =========PARAMS=========

	currentInteraction := &structures.UserInteraction{
		IsInteracting: true,
		Type:          structures.Commands["steam_topUpBalance"],
		Step:          0,
		DataCase:      make([]string, 2),
	}
	structures.UserStates[c.Chat().ID] = *currentInteraction

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
