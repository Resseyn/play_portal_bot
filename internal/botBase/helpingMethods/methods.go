package helpingMethods

import (
	"math/rand"
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
