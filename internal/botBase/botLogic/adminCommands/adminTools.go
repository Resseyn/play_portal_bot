package adminCommands

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"math/rand"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/databaseModels"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
	"strconv"
	"strings"
)

// мне было супервпадлу делать чтобы не нужно было создавать две отдельных функции
// но оказалось мне похуй и пусть все будет вот так, создание сообщения в botCommands.go
func ShowAdminPanel(c telebot.Context) error {

	// =========PARAMS=========
	picPath := "pkg/utils/data/img/adminImages/ramzes.jpg"
	messageContent := fmt.Sprintf("админ менюшка (тест %.2f)", rand.Float32()*1.3)
	messageData := &structures.MessageData{
		Command:     structures.Commands["adminPanel"],
		PrevCommand: "",
		Price:       0,
		Custom:      "",
	}
	commands := [][]structures.Command{
		{
			{Text: "Показать репорты", Command: structures.Commands["showReports"]},
			{Text: "Падрачить хуй", Command: structures.Commands["showAdminPanel"]},
			{Text: "Создать товар", Command: structures.Commands["createNewProduct"]}}}

	// =========PARAMS=========

	keyboard := helpingMethods.CreateInline(messageData, commands...)
	err := c.Edit(&telebot.Photo{
		File:    telebot.FromDisk(picPath),
		Caption: messageContent,
	}, keyboard)

	return err

}

// условно из FAQ идет то самое создание тикета и вот допустим они сюда идут ну я думаю потом придумаем куда это чисто временная прикольная хуйня
func ShowReports(c telebot.Context) error {
	// =========PARAMS=========
	picPath := "pkg/utils/data/img/adminImages/astronaut.jpg"
	messageContent := "долбаебы понапишут вопросов через факью и они все сюда"
	messageData := &structures.MessageData{
		Command:     structures.Commands["showReports"],
		PrevCommand: structures.Commands["showAdminPanel"],
		Price:       0,
		Custom:      "",
	}
	commands := [][]structures.Command{{}}
	// =========PARAMS=========

	keyboard := helpingMethods.CreateInline(messageData, commands...)
	err := c.Edit(&telebot.Photo{
		File:    telebot.FromDisk(picPath),
		Caption: messageContent,
	}, keyboard)
	return err
}

func CreateNewProduct(c telebot.Context) error {
	// =========PARAMS=========
	messageContent := "Отправь краткое название товара без пробелов на английском (spotify, appStore)"
	// =========PARAMS=========
	randCommand := helpingMethods.GenerateUniqueCommand(structures.Commands)
	structures.CreatingStates[c.Chat().ID] = &structures.CreatingState{Step: 0, EndStep: -1, MainCommand: randCommand, PrevPage: structures.Commands["shop_gameServices"]}
	err := c.Send(messageContent)
	return err
}

func HandleCreatingState(c telebot.Context) error {
	state := structures.CreatingStates[c.Chat().ID]
	if state.Step == 0 && c.Message().Text != "" {
		state.MainCommandName = c.Message().Text
		state.Step++
		c.Send("Теперь отправь картинку главной страницы товара")
	} else if state.Step == 1 && c.Message().Photo != nil {
		state.PicFIleID = c.Message().Photo.FileID
		state.Step++
		c.Send("Теперь отправь главный текст товара")
	} else if state.Step == 2 && c.Message().Text != "" {
		state.MainText = c.Message().Text
		state.Step++
		c.Send("Теперь отправь количество типов товара. Если 0 - создастся обычная страница")
	} else if state.Step == 3 && c.Message().Text != "" {
		var err error
		state.NumberOfGoods, err = strconv.Atoi(c.Message().Text)
		if err != nil {
			c.Send("цифрой")
			return err
		}
		if state.NumberOfGoods == 0 {
			state.Step = 6
			c.Send("Теперь скинь команду для страницы, с которой должна быть страница")
			return nil
		}
		state.Goods = make([]structures.Good, state.NumberOfGoods)
		state.Prices = make([]float64, state.NumberOfGoods)
		state.GoodCommands = make([]string, state.NumberOfGoods)
		state.GoodCommandsNames = make([]string, state.NumberOfGoods)
		state.GoodCustoms = make([]string, state.NumberOfGoods)
		state.Step++
		c.Send("Теперь отправь название хэндлера для товара. Если товар - ключ, напиши keyHandler")
	} else if state.Step == 4 && c.Message().Text != "" {
		if c.Message().Text == "keyHandler" {
			state.Step = 6
			c.Send("Теперь скинь команду для страницы, с которой должен быть товар")
			state.EndStep = 7 + (state.NumberOfGoods * 4) + 1
			return nil
		}
		state.Handler = c.Message().Text
		state.Step++
		c.Send("Теперь через запятую фразы для получения данных от юзера (Введи пароль, Введи логин)")
	} else if state.Step == 5 && c.Message().Text != "" {
		params := strings.Split(c.Message().Text, ",")
		state.HandlerParams = params
		state.Step++
		c.Send("Теперь скинь команду для страницы, с которой должен быть товар")

	} else if state.Step == 6 && c.Message().Text != "" {
		if _, ok := structures.Commands[c.Message().Text]; !ok {
			c.Send("Такой команды нет!")
			return nil
		}
		state.PrevPage = structures.Commands[c.Message().Text]
		if state.NumberOfGoods == 0 {
			state.Step = state.EndStep
			c.Send("Напиши еще что-то и создатся страница")
			return nil
		}
		state.Step++
		c.Send("Теперь заполним каждый товар. Скинь картинку первого")
		state.EndStep = 7 + (state.NumberOfGoods * 4)

	} else if state.Step > 6 && state.Step < state.EndStep {
		currentGood := (state.Step - 7) / 4
		switch (state.Step - 7) % 4 {
		case 0:
			if c.Message().Photo == nil {
				c.Send("пикчу")
			} else {
				randCommand := helpingMethods.GenerateUniqueCommand(structures.Commands)
				randCustom := helpingMethods.GenerateUniqueCommand(structures.Codes)
				state.Goods[currentGood] = structures.Good{Command: randCommand, Custom: randCustom}
				state.GoodCommands[currentGood] = state.Goods[currentGood].Command
				state.GoodCustoms[currentGood] = state.Goods[currentGood].Custom
				state.Goods[currentGood].URL = c.Message().Photo.FileID
				state.Step++
				c.Send("Теперь отправь главный текст товара")
			}
		case 1:
			if c.Message().Text == "" {
				c.Send("текст")
			} else {
				state.Goods[currentGood].Text = c.Message().Text
				state.Step++
				c.Send("Теперь отправь цену товара")
			}
		case 2:
			if c.Message().Text == "" {
				c.Send("цену")
			} else {
				price, err := strconv.ParseFloat(c.Message().Text, 64)
				if err != nil {
					c.Send("цену")
					return err
				}
				state.Prices[currentGood] = price
				state.Step++
				c.Send("Теперь отправь краткое название товара без пробелов (Спотифай индивидуал 1 месяц, АппСтр)")
			}
		case 3:
			if c.Message().Text == "" {
				c.Send("код")
			} else {
				state.GoodCommandsNames[currentGood] = c.Message().Text
				state.Step++
				if state.Step == state.EndStep {
					fmt.Println(*state)
					inlineCommands := make([][]structures.Command, state.NumberOfGoods)
					if len(state.GoodCommands) >= 8 {
						for i, command := range state.GoodCommands {
							comm := structures.Command{Text: state.GoodCommandsNames[i],
								Command: command}
							if i%2 == 0 {
								inlineCommands[i/2] = []structures.Command{comm}
							} else {
								inlineCommands[i/2] = append(inlineCommands[i/2], comm)
							}

						}
					} else {
						for i, command := range state.GoodCommands {
							comm := structures.Command{Text: state.GoodCommandsNames[i],
								Command: command}
							inlineCommands[i] = []structures.Command{comm}
						}
					}

					page := &structures.TypicalPage{
						URL:         state.PicFIleID,
						Text:        state.MainText,
						MainCommand: state.MainCommand,
						Commands:    inlineCommands,
						PrevPage:    state.PrevPage,
						Goods:       state.Goods,
					}
					handlerName := state.Handler
					handlerParams := state.HandlerParams
					codesText := make(map[string]string, state.NumberOfGoods)
					for i, name := range state.GoodCommandsNames {
						codesText[state.GoodCustoms[i]] = name
					}
					goodPages := make([]*structures.TypicalPage, state.NumberOfGoods)
					for i, _ := range state.Goods {
						goodPages[i] = page //TODO: у меня для каждого товара идет главная страница, потом оно магией распределяется, так что оптимизировать чуток надо
					}
					mainCommandName := state.MainCommandName
					err := databaseModels.AddNewPageToMongo(page, goodPages, mainCommandName, handlerName, handlerParams, state.Prices, codesText)
					if err != nil {
						loggers.ErrorLogger.Fatal(err)
						return err
					}
					c.Send("Товар добавлен в бд!")
					delete(structures.CreatingStates, c.Chat().ID)
					return nil
				}
				c.Send("Теперь картинку следующего товара")
			}
		}
	} else if state.Step == state.EndStep { //for normal pages
		fmt.Println(*state)
		inlineCommands := make([][]structures.Command, state.NumberOfGoods)

		page := &structures.TypicalPage{
			URL:         state.PicFIleID,
			Text:        state.MainText,
			MainCommand: state.MainCommand,
			Commands:    inlineCommands,
			PrevPage:    state.PrevPage,
		}
		handlerName := ""
		handlerParams := make([]string, 0)
		codesText := make(map[string]string, 0)
		goodPages := make([]*structures.TypicalPage, 0)
		mainCommandName := state.MainCommandName
		err := databaseModels.AddNewPageToMongo(page, goodPages, mainCommandName, handlerName, handlerParams, state.Prices, codesText)
		if err != nil {
			loggers.ErrorLogger.Fatal(err)
			return err
		}
		c.Send("Страница добавлена в бд!")
		delete(structures.CreatingStates, c.Chat().ID)
		return nil
	}
	return nil
}
