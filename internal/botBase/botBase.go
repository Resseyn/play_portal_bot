package botBase

import (
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/botBase/botCommands"
	"play_portal_bot/internal/botBase/botLogic"
	"play_portal_bot/internal/botBase/botLogic/mainMenuButtons"
	"play_portal_bot/internal/botBase/botLogic/mainMenuButtons/shopButtons"
	"play_portal_bot/internal/botBase/botLogic/mainMenuButtons/shopButtons/steamButtons"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/botBase/keys"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
	"time"
)

func BotStart() error {
	settings := telebot.Settings{
		Token:  keys.BotKey,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := telebot.NewBot(settings)
	if err != nil {
		loggers.ErrorLogger.Fatal(err)
		return err
	}

	b.Handle("/start", botCommands.Start)
	b.Handle(telebot.OnCallback, CallbackHandle)
	b.Handle(telebot.OnText, func(c telebot.Context) error {
		if _, ok := structures.UserStates[c.Chat().ID]; ok {
			if structures.UserStates[c.Chat().ID].IsInteracting {
				return nil
			} else {
				return botCommands.Start(c)
			}
		} else {
			return botCommands.Start(c)
		}
	})
	b.Start()
	return nil
}

// CallbackHandle handles all the existing callbacks
func CallbackHandle(c telebot.Context) error {
	data := helpingMethods.ParseData(c.Callback().Data)
	switch data.Command {
	case "mainMenu":
		return botLogic.Menu(c)
	case "shop":
		return botLogic.Shop(c)
	case "personalCabinet":
		return botLogic.PersonalCabinet(c)
	case "support":
		return botLogic.Support(c)
	case "faq":
		return botLogic.FAQ(c)
	case "shop_gameServices": //TODO: change to shop_gameServices
		return mainMenuButtons.GameServices(c)
	case "shop_services":
		return mainMenuButtons.Services(c)
	case "shop_gameServices_steam":
		return shopButtons.Steam(c)

	case "steam_topUpBalance":
		return steamButtons.SteamTopUpBalance(c)
	}
	return nil
}
