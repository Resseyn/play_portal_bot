package helpingMethods

import (
	"gopkg.in/telebot.v3"
	"math/rand"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
	"strconv"
	"strings"
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
	data := ParseData(c.Callback().Data)
	params := structures.Pages[data.Command]
	if params.Data != nil {
		data = params.Data
	}
	data.PrevCommand = params.PrevPage
	data.Custom = params.Custom
	data.Price = int(structures.Prices[data.Custom])
	msg := &telebot.Photo{
		File:    telebot.FromDisk(params.URL),
		Caption: params.Text,
	}
	keyboard := CreateInline(data, params.Commands...)
	err := c.Edit(msg, keyboard)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return err
	}
	return nil
}
