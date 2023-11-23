package botBase

import (
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/botBase/botCommands"
	"play_portal_bot/internal/botBase/botLogic"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/loggers"
	"time"
)

func BotStart() error {
	settings := telebot.Settings{
		Token:  BotKey,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(settings)
	if err != nil {
		loggers.ErrorLogger.Fatal(err)
		return err
	}

	b.Handle("/start", botCommands.Start)
	b.Handle(telebot.OnCallback, CallbackHandle)

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
	case "gameServices":
		return botLogic.GameServices(c)
	case "services":
		return botLogic.Services(c)
	}
	return nil
}
