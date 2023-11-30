package test

//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"log"
//	"net/http"
//	"play_portal_bot/internal/botBase/helpingMethods"
//	"play_portal_bot/internal/botBase/keys"
//	"play_portal_bot/internal/loggers"
//	"play_portal_bot/pkg/utils/structures"
//	"strconv"
//	"testing"
//)
//
//func TestAPI(t *testing.T) {
//	GOTTENFROMDBCHATID := 2038902313
//	chatID := strconv.Itoa(GOTTENFROMDBCHATID)
//	order := structures.Commands["spotifySuccess"]
//	url2 := fmt.Sprintf("https://api.telegram.org/bot%s/%s", keys.BotKey, "sendMessage")
//	msgData := &structures.MessageData{
//		Command:     "", //wtf
//		PrevCommand: "",
//	}
//	commands := [][]structures.Command{{
//		{Text: "Вернуться к услуге", Command: order}}}
//	keyboard := helpingMethods.CreateInline(msgData, commands...)
//	jsonKeyboard, err := json.Marshal(keyboard)
//	if err != nil {
//		loggers.ErrorLogger.Println(err)
//		return
//	}
//	data := map[string]string{"chat_id": chatID, "text": fmt.Sprintf("hallo frend i got ya maney for %v", order), "reply_markup": string(jsonKeyboard)}
//	jsonData, err := json.Marshal(data)
//	if err != nil {
//		loggers.ErrorLogger.Println(err)
//		return
//	}
//	//params := url.Values{}
//	//params.Add("chat_id", chatID)
//	//params.Add("text", "hallo frend i got ya maney")
//	//st := params.Encode()
//	req, err := http.NewRequest("POST", url2, bytes.NewBuffer(jsonData))
//	req.Header.Add("Content-Type", "application/json")
//	client := &http.Client{}
//	answer, err := client.Do(req)
//	if err != nil {
//		log.Fatal(err)
//		return
//	}
//	defer answer.Body.Close()
//	var str []byte
//	answer.Body.Read(str)
//	structures.UserStates[int64(GOTTENFROMDBCHATID)] = &structures.UserInteraction{
//		IsInteracting: true,
//		Type:          order,
//		Step:          0,
//		DataCase:      make([]string, 2),
//	}
//	t.Log(string(str))
//}
