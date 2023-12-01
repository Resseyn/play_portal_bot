package helpingMethods

import (
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
