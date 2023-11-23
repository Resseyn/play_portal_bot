package botBase

import (
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/botBase/botCommands"
	"play_portal_bot/internal/botBase/botLogic"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/loggers"
	"time"
)

func BotStart() {
	settings := telebot.Settings{
		Token:  BotKey,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(settings)
	if err != nil {
		loggers.ErrorLogger.Fatal(err)
		return
	}

	b.Handle("/start", botCommands.Start)
	b.Handle(telebot.OnCallback, CallbackHandle)

	b.Start()
}
func CallbackHandle(c telebot.Context) error {
	data := helpingMethods.ParseData(c.Callback().Data)
	switch data.Command {
	case "mainMenu":
		return botLogic.Menu(c)
	case "showShop":
		return botLogic.ShowShop(c)
	}
	return nil
}
