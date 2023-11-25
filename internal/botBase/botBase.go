package botBase

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/botBase/botCommands"
	"play_portal_bot/internal/botBase/botLogic"
	"play_portal_bot/internal/botBase/botLogic/mainMenuButtons"
	"play_portal_bot/internal/botBase/botLogic/mainMenuButtons/shopButtons"
	"play_portal_bot/internal/botBase/botLogic/mainMenuButtons/shopButtons/servicesButtons"
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
	//b.Handle(telebot.OnCheckout, func(c telebot.Context) error {
	//	ans := answer{Id: c.Update().PreCheckoutQuery.ID, Ok: true, ErrorMessage: ""}
	//	resp, err := ans.Send(c.Bot(), telebot.ChatID(1), &telebot.SendOptions{})
	//	fmt.Println(resp.Payment)
	//	if err != nil {
	//		loggers.ErrorLogger.Println(err)
	//		return err
	//	}
	//	c.Send("аххаха заскамлен)))")
	//	return nil
	//})
	b.Start()
	return nil
}

// CallbackHandle handles all the existing callbacks
func CallbackHandle(c telebot.Context) error {
	fmt.Println(c.Callback().Data)
	data := helpingMethods.ParseData(c.Callback().Data)
	switch data.Command {
	case structures.Commands["buy"]:
		return helpingMethods.TopUpBalance(c)
	case structures.Commands["createCheck"]:
		return helpingMethods.CreateBill(c)
	case structures.Commands["mainMenu"]:
		return botLogic.Menu(c)
	case structures.Commands["shop"]:
		return botLogic.Shop(c)
	case structures.Commands["personalCabinet"]:
		return botLogic.PersonalCabinet(c)
	case structures.Commands["support"]:
		return botLogic.Support(c)
	case structures.Commands["faq"]:
		return botLogic.FAQ(c)
	case structures.Commands["shop_gameServices"]: //TODO: change to shop_gameServices
		return mainMenuButtons.GameServices(c)
	case structures.Commands["shop_services"]:
		return mainMenuButtons.Services(c)
	case structures.Commands["shop_gameServices_steam"]:
		return shopButtons.Steam(c)
	case structures.Commands["spotify"]:
		return servicesButtons.Spotify(c)
	case structures.Commands["spotify_individual_1"]:
		return servicesButtons.Spotify_Individual_1(c)
	case structures.Commands["steam_topUpBalance"]:
		return steamButtons.SteamTopUpBalance(c)
	}
	return nil
}

//type answer struct {
//	Id           string `json:"pre_checkout_query_id"`
//	Ok           bool   `json:"ok"`
//	ErrorMessage string `json:"error_message"`
//	telebot.Sendable
//}
//
//func (i *answer) Send(b *telebot.Bot, to telebot.Recipient, opt *telebot.SendOptions) (*telebot.Message, error) {
//	params := make(map[string]string)
//	params["pre_checkout_query_id"] = i.Id
//	params["ok"] = strconv.FormatBool(i.Ok)
//	params["error_message"] = i.ErrorMessage
//
//	data, err := b.Raw("answerPreCheckoutQuery", params)
//	if err != nil {
//		return nil, err
//	}
//	return extractMessage(data)
//}
//func extractMessage(data []byte) (*telebot.Message, error) {
//	var resp struct {
//		Result *telebot.Message
//	}
//	if err := json.Unmarshal(data, &resp); err != nil {
//		var resp struct {
//			Result bool
//		}
//		if err := json.Unmarshal(data, &resp); err != nil {
//			return nil, err
//		}
//		if resp.Result {
//			return nil, telebot.ErrTrueResult
//		}
//		return nil, err
//	}
//	return resp.Result, nil
//}
