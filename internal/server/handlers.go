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

	truePayment, err := onlineCasses.PayoutStatus(orderID)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		http.Error(w, "wrong method", http.StatusBadRequest)
		return
	}

	status := r.Form.Get("Status")
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

		url2 := fmt.Sprintf("https://api.telegram.org/bot%s/%s", keys.BotKey, "sendMessage")
		msgData := &structures.MessageData{
			Command:     "",
			PrevCommand: "",
			Custom:      order.Custom,
			Price:       int(structures.Prices[order.Custom]),
		}
		var commands [][]structures.Command
		if order.Custom == "aaac" {
			commands = [][]structures.Command{{
				{Text: "Успешная оплата", Command: ""},
			}}
		} else {
			commands = [][]structures.Command{{
				{Text: "Вернуться к услуге", Command: order.Custom}}}
		}
		keyboard := helpingMethods.CreateInline(msgData, commands...)
		jsonKeyboard, err := json.Marshal(keyboard)
		if err != nil {
			loggers.ErrorLogger.Println(err)
			return
		}
		data := map[string]string{"chat_id": strconv.FormatInt(order.ChatID, 10), "text": fmt.Sprintf("hallo frend i got ya maney for %v", order), "reply_markup": string(jsonKeyboard)}
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
		//TODO:send to admin db payment status from PayoutStatus func like "order_id" , "account_id", "chat_id"(opt, get from связывания таблиц), "status"
		//TODO: create func to return all payments with SUCCESS status (for admin)
	} else {
		//TODO:иди нахуй черт
	}
}

// PayPalychSuccessPaymentHandler метод для обработки постбек после успешной оплаты, смотря на кастом проводится услуга
func PayPalychSuccessPaymentHandler(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//defer r.Body.Close()
	//
	//OrderID := r.Form.Get("InvId")
	////TODO:db search OrderID and return order (orderID и accountID, а так же услуга и ее цена), а так же считать заказ выполненным
	////TODO:db search если пополненная сумма больше или равна нужной для оплаты услуги, продолжить код, иначе чет другое типо иди нахуй
	//fmt.Println(OrderID)
	//GOTTENFROMDBCHATID := 2038902313
	//chatID := strconv.Itoa(GOTTENFROMDBCHATID)
	//command := structures.Commands[r.Form.Get("mainMenu")]
	//url2 := fmt.Sprintf("https://api.telegram.org/bot%s/%s", keys.BotKey, "sendMessage")
	//msgData := &structures.MessageData{
	//	Command:     "", //wtf
	//	PrevCommand: "",
	//}
	//commands := [][]structures.Command{{
	//	{Text: "Главное меню", Command: command}}}
	//keyboard := helpingMethods.CreateInline(msgData, commands...)
	//jsonKeyboard, err := json.Marshal(keyboard)
	//if err != nil {
	//	loggers.ErrorLogger.Println(err)
	//	return
	//}
	//data := map[string]string{"chat_id": chatID, "text": fmt.Sprintf("hallo frend i got ya maney, chill"), "reply_markup": string(jsonKeyboard)}
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
	//var str []byte
	//answer.Body.Read(str)
	//w.Write(str)
	//TODO: просто редирект на страницу с подписью что оплата прошла все ок ващеее
}

func PayPalychFailPaymentHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	defer r.Body.Close()

	OrderID := r.Form.Get("InvId")
	//TODO:db search OrderID and return order (orderID и accountID, а так же услуга и ее цена), а так же считать заказ выполненным
	//TODO:db search если пополненная сумма больше или равна нужной для оплаты услуги, продолжить код, иначе чет другое типо иди нахуй
	fmt.Println(OrderID)
	GOTTENFROMDBCHATID := 2038902313
	chatID := strconv.Itoa(GOTTENFROMDBCHATID)
	command := structures.Commands[r.Form.Get("mainMenu")]
	url2 := fmt.Sprintf("https://api.telegram.org/bot%s/%s", keys.BotKey, "sendMessage")
	msgData := &structures.MessageData{
		Command:     "", //wtf
		PrevCommand: "",
	}
	commands := [][]structures.Command{{
		{Text: "Главное меню", Command: command}}}
	keyboard := helpingMethods.CreateInline(msgData, commands...)
	jsonKeyboard, err := json.Marshal(keyboard)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return
	}
	data := map[string]string{"chat_id": chatID, "text": fmt.Sprintf("hallo suka i didnt get ya maney, trai agen mthf"), "reply_markup": string(jsonKeyboard)}
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
	var str []byte
	answer.Body.Read(str)
	w.Write(str)
}
