package botBase

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"play_portal_bot/internal/botBase/botCommands"
	"play_portal_bot/internal/botBase/botLogic"
	"play_portal_bot/internal/botBase/botLogic/adminCommands"
	"play_portal_bot/internal/botBase/botLogic/mainMenuButtons"
	"play_portal_bot/internal/botBase/botLogic/mainMenuButtons/shopButtons"
	"play_portal_bot/internal/botBase/botLogic/mainMenuButtons/shopButtons/servicesButtons"
	"play_portal_bot/internal/botBase/botLogic/mainMenuButtons/shopButtons/steamButtons"
	"play_portal_bot/internal/botBase/botLogic/mainMenuButtons/supportMethods"
	"play_portal_bot/internal/botBase/botLogic/orderMethods"
	"play_portal_bot/internal/botBase/botLogic/successfulPayments"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/botBase/keys"
	"play_portal_bot/internal/botBase/onlineCasses"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
	"strconv"
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
	b.Handle("/end", supportMethods.EndTicket)
	b.Handle("/endOrder", orderMethods.EndOrder)
	b.Handle(fmt.Sprintf("/tool%v", keys.ToolKey), botCommands.CreateAdminPanel)

	b.Handle(telebot.OnCallback, CallbackHandle)

	b.Handle(telebot.OnText, func(c telebot.Context) error {
		if state, ok := structures.UserStates[c.Chat().ID]; ok {
			fmt.Println("STATE:", state)
			if state.IsInteracting {
				switch structures.UserStates[c.Chat().ID].Type {

				case "moderatorDialog":
					convWith, _ := strconv.Atoi(state.DataCase[0])
					// if в случае если юзер вышел нахуй из диалога с модером а модер не договорил
					if _, ok := structures.UserStates[int64(convWith)]; !ok {
						structures.UserStates[int64(convWith)] = &structures.UserInteraction{
							IsInteracting: true,
							Type:          "moderatorDialog",
							DataCase:      []string{strconv.FormatInt(c.Chat().ID, 10)},
						}
					}
					// if в случае если юзер вышел нахуй из диалога с модером а модер не договорил
					c.Bot().Send(telebot.ChatID(convWith), c.Message().Text)

				case "awaitingForPrice":
					newPrice, err := strconv.ParseFloat(c.Message().Text, 64)
					if err != nil {
						commands := [][]structures.Command{
							{{Text: "Отмена", Command: structures.Commands["mainMenu"]}},
						}
						keyboard := helpingMethods.CreateInline(&structures.MessageData{}, commands...)
						c.Send("Пожалуйста, введите цифру", keyboard)
						return err
					}
					state.Price = newPrice
					return orderMethods.CreateCheck(c)

				case "spotifyHandler":
					state.DataCase[state.Step] = c.Message().Text
					if state.Step == 1 {
						return orderMethods.CreateOrder(c)
					} else {
						state.Step++
						return sucessfulPayments.SpotifySuccessPayment(c)
					}

				default:
					return botCommands.Start(c)
				}
			} else {
				return botCommands.Start(c)
			}
		} else {
			return botCommands.Start(c)
		}
		return nil
	})
	b.Start()
	return nil
}

// CallbackHandle handles all the existing callbacks
func CallbackHandle(c telebot.Context) error {
	data := helpingMethods.ParseData(c.Callback().Data)
	fmt.Println("STATE:", structures.UserStates[c.Chat().ID])
	fmt.Println("CALLBACK DATA:", data)
	switch data.Command {
	case structures.Commands["topUpBalance"]:
		return orderMethods.TopUpBalance(c)
	case structures.Commands["createCheck"]:
		return orderMethods.CreateCheck(c)
	case structures.Commands["createPayPalychBill"]:
		return onlineCasses.CreatePayPalychBill(c)
	case structures.Commands["respondToOrder"]:
		return orderMethods.RespondToOrder(c)
	case structures.Commands["endOrder"]:
		return orderMethods.EndOrder(c)
	//from mainMenu==============================
	case structures.Commands["mainMenu"]:
		return botLogic.Menu(c)
	case structures.Commands["shop"]:
		return botLogic.Shop(c)

	//from menu_shop===============================
	case structures.Commands["shop_gameServices"]:
		return mainMenuButtons.GameServices(c)
	case structures.Commands["shop_services"]:
		return mainMenuButtons.Services(c)
	case structures.Commands["shop_gameServices_steam"]:
		return shopButtons.Steam(c)

	case structures.Commands["spotify"]:
		return servicesButtons.Spotify(c)
	case structures.Commands["spotify_individual_1"]:
		return servicesButtons.SpotifyIndividual1(c)

	case structures.Commands["spotifySuccessIND1"]:
		return sucessfulPayments.SpotifySuccessPayment(c)
	case structures.Commands["spotifySuccessIND3"]:
		return sucessfulPayments.SpotifySuccessPayment(c)
	case structures.Commands["spotifySuccessIND6"]:
		return sucessfulPayments.SpotifySuccessPayment(c)
	case structures.Commands["spotifySuccessIND12"]:
		return sucessfulPayments.SpotifySuccessPayment(c)
	case structures.Commands["spotifySuccessDUO1"]:
		return sucessfulPayments.SpotifySuccessPayment(c)
	case structures.Commands["spotifySuccessFAM1"]:
		return sucessfulPayments.SpotifySuccessPayment(c)

	case structures.Commands["steam_topUpBalance"]:
		return steamButtons.SteamTopUpBalance(c)
	//============================================

	case structures.Commands["personalCabinet"]:
		return botLogic.PersonalCabinet(c)
	case structures.Commands["support"]:
		return botLogic.Support(c)

	//from menu_support===========================
	case structures.Commands["createTicket"]:
		return supportMethods.CreateTicket(c)
	case structures.Commands["respondToTicket"]:
		return supportMethods.RespondToTicket(c)
	case structures.Commands["endTicket"]:
		return supportMethods.EndTicket(c)
	//============================================

	case structures.Commands["faq"]:
		return botLogic.FAQ(c)

	//from adminPanel=============================
	case structures.Commands["showAdminPanel"]:
		return adminCommands.ShowAdminPanel(c)
	case structures.Commands["showReports"]:
		return adminCommands.ShowReports(c)

	case structures.Commands["pingModer"]:
		return orderMethods.PingModer(c)
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
