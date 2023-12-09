package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"play_portal_bot/internal/botBase/helpingMethods"
	"play_portal_bot/internal/botBase/keys"
	"play_portal_bot/internal/botBase/onlineCasses"
	"play_portal_bot/internal/databaseModels"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/utils/structures"
	"strconv"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("web/html/home.page.tmpl.html")
	if err != nil {
		loggers.ErrorLogger.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

// PayPalychPaymentHandler метод для обработки постбек после успешной оплаты, смотря на кастом проводится услуга
func PayPalychPaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "wrong method", http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	orderID := r.Form.Get("InvId")

	//truePayment, err := onlineCasses.PayoutStatus(orderID) TODO: posle moderki aga
	//if err != nil {
	//	loggers.ErrorLogger.Println(err)
	//	http.Error(w, "wrong method", http.StatusBadRequest)
	//	return
	//}//TODO: вообще все онли с пэйпалыча брать
	truePayment := onlineCasses.Payment{Status: "SUCCESS"} //TODO: сравнивать еще и сумму

	status := r.Form.Get("Status")
	//signature := r.Form.Get("SignatureValue")
	outSum, _ := strconv.ParseFloat(r.Form.Get("OutSum"), 64)
	commision, _ := strconv.ParseFloat(r.Form.Get("Commission"), 64)
	amount := outSum - commision
	order, err := databaseModels.Orders.GetOrder(orderID)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		http.Error(w, "OrderNotFound", 404)
		return
	}
	if status == "SUCCESS" && status == truePayment.Status && order.Status != status {

		_, err = databaseModels.Orders.OrderIsDone(orderID)
		if err != nil {
			loggers.ErrorLogger.Println(err)
			http.Error(w, "OrderNotFound", http.StatusBadRequest)
			return
		}

		_, err = databaseModels.Users.TopUpBalance(order.ChatID, amount)
		if err != nil {
			loggers.ErrorLogger.Println(err)
			http.Error(w, "UserNotFound, how?", 404)
			return
		}
		//_, err = databaseModels.Orders.CreateCheck(order.OrderID, order.ChatID, amount, "aaac")

		var commands [][]structures.Command
		var msgData *structures.MessageData

		if _, ok := structures.UserOrders[order.ChatID]; !ok {
			commands = [][]structures.Command{{
				{Text: "Успешная оплата", Command: structures.Commands["mainMenu"]},
			}}
			msgData = &structures.MessageData{}
		} else {
			commands = [][]structures.Command{{
				{Text: "Вернуться к услуге", Command: structures.Commands["spotifySuccess"]}}}
			msgData = &structures.MessageData{
				Command:     "",
				PrevCommand: "",
				Custom:      structures.UserOrders[order.ChatID],
				Price:       int(structures.Prices[structures.UserOrders[order.ChatID]]),
			}
		}

		url2 := fmt.Sprintf("https://api.telegram.org/bot%s/%s", keys.BotKey, "sendMessage")

		keyboard := helpingMethods.CreateInline(msgData, commands...)
		jsonKeyboard, err := json.Marshal(keyboard)
		if err != nil {
			loggers.ErrorLogger.Println(err)
			return
		}
		data := map[string]string{"chat_id": strconv.FormatInt(order.ChatID, 10), "text": "Оплата прошла. Удачных покупок!", "reply_markup": string(jsonKeyboard)}
		jsonData, err := json.Marshal(data)
		if err != nil {
			loggers.ErrorLogger.Println(err)
			return
		}
		req, err := http.NewRequest("POST", url2, bytes.NewBuffer(jsonData))
		req.Header.Add("Content-Type", "application/json")
		client := &http.Client{}
		answer, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer answer.Body.Close()

		//TODO: create func to return all payments with SUCCESS status (for admin)
	} else {
		url2 := fmt.Sprintf("https://api.telegram.org/bot%s/%s", keys.BotKey, "sendMessage")
		data := map[string]string{"chat_id": strconv.FormatInt(order.ChatID, 10), "text": fmt.Sprintf("Оплата по заказу %v не прошла. Пожалуйста, попробуйте ещё раз или обратитесь в службу поддержки платежа", orderID)}
		jsonData, err := json.Marshal(data)
		if err != nil {
			loggers.ErrorLogger.Println(err)
			return
		}
		req, err := http.NewRequest("POST", url2, bytes.NewBuffer(jsonData))
		req.Header.Add("Content-Type", "application/json")
		client := &http.Client{}
		answer, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer answer.Body.Close()
	}
}

// PayPalychSuccessPaymentHandler метод для обработки постбек после успешной оплаты, смотря на кастом проводится услуга
func PayPalychSuccessPaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "wrong method", http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()

	//truePayment, err := onlineCasses.PayoutStatus(orderID) TODO: posle moderki aga
	//if err != nil {
	//	loggers.ErrorLogger.Println(err)
	//	http.Error(w, "wrong method", http.StatusBadRequest)
	//	return
	//}//TODO: вообще все онли с пэйпалыча брать
	truePayment := onlineCasses.Payment{Status: "SUCCESS"} //TODO: сравнивать еще и сумму и signatureValue

	//signature := r.Form.Get("SignatureValue") //TODO: это просто метод который делает редирект чзх
	if truePayment.Status == "SUCCESS" {
		http.Redirect(w, r, "https://t.me/play_portal_bot", 200)
	}

}

func PayPalychFailPaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "wrong method", http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	orderID := r.Form.Get("InvId")

	//truePayment, err := onlineCasses.PayoutStatus(orderID) TODO: posle moderki aga
	//if err != nil {
	//	loggers.ErrorLogger.Println(err)
	//	http.Error(w, "wrong method", http.StatusBadRequest)
	//	return
	//}//TODO: вообще все онли с пэйпалыча брать
	truePayment := onlineCasses.Payment{Status: "FAIL"} //TODO: сравнивать еще и сумму и signatureValue

	status := r.Form.Get("Status")
	//signature := r.Form.Get("SignatureValue")
	order, err := databaseModels.Orders.GetOrder(orderID)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		http.Error(w, "OrderNotFound", 404)
		return
	}
	if status == truePayment.Status && order.Status != status {
		//url2 := fmt.Sprintf("https://api.telegram.org/bot%s/%s", keys.BotKey, "sendMessage")
		//data := map[string]string{"chat_id": strconv.FormatInt(order.ChatID, 10), "text": fmt.Sprintf("Оплата по заказу %v не прошла. Пожалуйста, попробуйте ещё раз или обратитесь в службу поддержки платежа", orderID)}
		//jsonData, err := json.Marshal(data)
		//if err != nil {
		//	loggers.ErrorLogger.Println(err)
		//	return
		//}
		//req, err := http.NewRequest("POST", url2, bytes.NewBuffer(jsonData))
		//req.Header.Add("Content-Type", "application/json")
		//client := &http.Client{}
		//answer, err := client.Do(req)
		//if err != nil {
		//	log.Fatal(err)
		//	return
		//}
		//defer answer.Body.Close()
	}
	http.Redirect(w, r, "https://t.me/play_portal_bot", 200)
}
