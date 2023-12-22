package helpingMethods

import (
	"encoding/json"
	"fmt"
	"gopkg.in/telebot.v3"
	"math/rand"
	"os"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
	"strconv"
	"strings"
	"time"
)

// ParseData парсит строку с колбек кновпеи в структуру
func ParseData(callbackData string) *structures.MessageData {
	callbackData = strings.Trim(callbackData, "\n")
	data := strings.Split(callbackData, ",") //0 - chatID 1- messageID 2 - command 3 - prevCommand
	data2, _ := strconv.Atoi(data[2])
	messageData := &structures.MessageData{
		Command:     data[0],
		PrevCommand: data[1],
		Price:       data2,
		Custom:      data[3],
	}
	return messageData
}

// NewInteraction creates new interaciton for user, optPrice and optData is optional
func NewInteraction(interactionType string, chatID int64, optPrice float64, optOrder string, optData []string) {
	delete(structures.UserStates, chatID)
	structures.UserStates[chatID] = &structures.UserInteraction{
		IsInteracting: true,
		Type:          interactionType,
		Step:          0,
		Price:         optPrice,
		Order:         optOrder,
		DataCase:      optData,
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandStringRunes создает рандомный OrderID
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func FindKeyByValue(m map[string]string, value string) (key string, ok bool) {
	for k, v := range m {
		if v == value {
			return k, true
		}
	}
	return "", false
}

func SendTypicalPage(c telebot.Context) error {
	fm := time.Now()
	data := ParseData(c.Callback().Data)
	params := structures.Pages[data.Command]
	data.PrevCommand = params.PrevPage
	if params.Goods != nil {
		if params.MainCommand != data.Command {
			if params.Data != nil {
				data = params.Data
			}
			for _, good := range params.Goods {
				if good.Command == data.Command {
					data.Custom = good.Custom
					data.Price = int(structures.Prices[good.Custom])
					data.PrevCommand = params.MainCommand
					var File telebot.File
					if len(good.URL) == 83 && good.URL[0:3] != "pkg" {
						File = telebot.File{FileID: good.URL}
					} else {
						File = telebot.FromDisk(good.URL)
					}
					msg := &telebot.Photo{
						File:    File,
						Caption: good.Text,
					}
					commands := [][]structures.Command{
						{
							{Text: "Купить", Command: structures.Commands["topUpBalance"]}},
					}
					keyboard := CreateInline(data, commands...)
					err := c.Edit(msg, keyboard)
					if err != nil {
						loggers.ErrorLogger.Println(err)
						return err
					}
					fmt.Println(time.Now().Sub(fm).Seconds())
					return nil
				}
			}
		}
	} else {
		if params.Data != nil {
			data = params.Data
			data.PrevCommand = params.PrevPage
		}
	}
	var File telebot.File
	if len(params.URL) == 83 && params.URL[0:3] != "pkg" {
		File = telebot.File{FileID: params.URL}
	} else {
		File = telebot.FromDisk(params.URL)
	}
	msg := &telebot.Photo{
		File:    File,
		Caption: params.Text,
	}
	keyboard := CreateInline(data, params.Commands...)
	err := c.Edit(msg, keyboard)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return err
	}
	fmt.Println(time.Now().Sub(fm).Seconds())
	return nil
}
func ParseMaps() {
	com, _ := json.Marshal(structures.Commands)
	os.WriteFile("jsons/commands.json", com, 0666)
	com, _ = json.Marshal(structures.Codes)
	os.WriteFile("jsons/codes.json", com, 0666)
	com, _ = json.Marshal(structures.Handlers)
	os.WriteFile("jsons/handlers.json", com, 0666)
	com, _ = json.Marshal(structures.Pages)
	os.WriteFile("jsons/pages.json", com, 0666)
	com, _ = json.Marshal(structures.Prices)
	os.WriteFile("jsons/prices.json", com, 0666)
	com, _ = json.Marshal(structures.Parameters)
	os.WriteFile("jsons/parameters.json", com, 0666)
}

func GenerateUniqueCommand(mapp map[string]string) string {
	randMainCommand := RandStringRunes(4)
	if _, ok := FindKeyByValue(mapp, randMainCommand); ok {
		return GenerateUniqueCommand(mapp) // Если ключ уже существует, вызываем функцию снова
	}
	return randMainCommand // Если ключа нет, возвращаем сгенерированную строку
}
