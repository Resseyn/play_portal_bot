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
	data0, _ := strconv.Atoi(data[0])
	data1, _ := strconv.Atoi(data[1])
	data4, _ := strconv.Atoi(data[4])
	messageData := &structures.MessageData{
		ChatID:      int64(data0),
		MessageID:   data1,
		Command:     data[2],
		PrevCommand: data[3],
		Price:       data4,
	}
	return messageData
}
