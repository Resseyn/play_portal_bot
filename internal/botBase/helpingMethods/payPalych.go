package helpingMethods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/telebot.v3"
	"io/ioutil"
	"log"
	"net/http"
	"play_portal_bot/internal/botBase/keys"
	"play_portal_bot/internal/loggers"
	"strconv"
)

type Bill struct {
	Amount              float64 `json:"amount"`
	OrderId             string  `json:"order_id"`
	Description         string  `json:"description"`
	Type                string  `json:"type"` //normal or multi
	ShopId              string  `json:"shop_id"`
	CurrencyIn          string  `json:"currency_in"` //RUB USD EUR
	Custom              string  `json:"custom"`
	PayerPaysCommission int     `json:"payer_pays_commission"` //1 или 0	Параметр, который указывает на то, кто будет оплачивать комиссию за входящий платёж.
	Name                string  `json:"name"`                  //Донат	любая строка	Название ссылки. Укажите, за что принимаете средства. Этот текст будет отображен в платежной форме.
	SuccessUrl          string  `json:"success_url"`           //https://exmpl/order_321/success	Страница успешной оплаты.
	FailUrl             string  `json:"fail_url"`              //https://exmpl/order_321/fail
}

// Payment what goes on RESULT URL application/x-www-form-urlencoded
type Payment struct {
	InvId           string  // Уникальный идентификатор заказа, переданный при формировании счета
	OutSum          float64 // Сумма платежа
	Commission      float64 // Комиссия с платежа
	TrsId           string  // Уникальный идентификатор платежа
	Status          string  // Статус платежа
	CurrencyIn      string  // Валюта, в которой оплачивался счет
	Custom          string  // Произвольное поле, переданное при формировании счета
	AccountType     string  // Метод оплаты
	AccountNumber   string  // Дополнительная информация о методе оплаты
	BalanceAmount   float64 // Сумма, которая зачислена на баланс
	BalanceCurrency string  // Валюта, в которой было зачисление денежных средств на баланс
	ErrorCode       string  // Код ошибки
	ErrorMessage    string  // Описание ошибки
	SignatureValue  string  // Подпись
}

type BillCreateResponse struct {
	Success     bool   `json:"success"`       //true/false	true	Флаг успешности запроса
	LinkUrl     string `json:"link_url"`      //URL	 	https://paypalych.com/link/5QWlqB2kKJ	Ссылка на страницу с QR кодом
	LinkPageUrl string `json:"link_page_url"` //URL	 	https://paypalych.com/transfer/5QWlqB2kKJ	Ссылка на оплату
	BillId      string `json:"bill_id"`       //5QWlqB2kKJ
}

func CreateBill(c telebot.Context) error {
	// Создание счета на PayPalych
	data := &Bill{
		Amount:              16.00,
		OrderId:             "123",
		Description:         "Описание ссылки",
		Type:                "normal",
		ShopId:              "G1vrEyX0LR",
		CurrencyIn:          "RUB",
		Custom:              "кастомное поле со свободным текстом",
		PayerPaysCommission: 1,
		Name:                "Платёж",
		SuccessUrl:          "https://t.me/play_portal_bot",
	}
	jsonData, _ := json.Marshal(*data)

	req, _ := http.NewRequest("POST", "https://paypalych.com/api/v1/bill/create", bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "Bearer 123|q4uNcWNKMNZoSFSY1XTxp36nsM0kUMSu0otSA95")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	var billCreateResponse BillCreateResponse
	err = json.Unmarshal(body, &billCreateResponse)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return err
	}

	// Отправка ссылки на оплату пользователю
	_, err = c.Bot().Send(telebot.ChatID(c.Chat().ID), "Пожалуйста, оплатите счет по этой ссылке: "+billCreateResponse.LinkPageUrl)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return err
	}
	return nil
}

type Payout struct {
	ID                string  `json:"id"`
	Status            string  `json:"status"`
	Amount            float64 `json:"amount"`
	AccountAmount     float64 `json:"account_amount"`
	Commission        float64 `json:"commission"`
	AccountIdentifier string  `json:"account_identifier"`
	Currency          string  `json:"currency"`
	AccountCurrency   string  `json:"account_currency"`
	ErrorCode         int     `json:"error_code"`
	CreatedAt         string  `json:"created_at"`
	Success           bool    `json:"success"`
}
type StatusParams struct {
	Id      string `json:"id"`
	OrderId string `json:"order_id"`
}

func PayoutStatus(c telebot.Context) error {
	data := &StatusParams{
		OrderId: "123",
	}
	jsonData, _ := json.Marshal(*data)
	req, _ := http.NewRequest("POST", "https://paypalych.com/api/v1/payout/status", bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "Bearer 123|q4uNcWNKMNZoSFSY1XTxp36nsM0kUMSu0otSA95")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	var billCreateResponse Payout
	err = json.Unmarshal(body, &billCreateResponse)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return err
	}

	_, err = c.Bot().Send(telebot.ChatID(c.Chat().ID), "status: "+billCreateResponse.Status)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return err
	}
	return nil
}

func PaymentHandle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//var payment Payment
	//defer r.Body.Close()
	//body, _ := io.ReadAll(r.Body)
	//
	//err := json.Unmarshal(body, &payment)
	//if err != nil {
	//	loggers.ErrorLogger.Println(err)
	//	return
	//}
	//if payment.Status == "SUCCESS"
	AccountID := r.Form.Get("Custom")
	//TODO:db search AccountID and return via chatID and UserStates to bot the TOVAR
	//TODO:db get from ChatID from AccountID
	fmt.Println(AccountID)
	GOTTENFROMDBCHATID := 2038902313
	chatID := strconv.Itoa(GOTTENFROMDBCHATID)
	url2 := fmt.Sprintf("https://api.telegram.org/bot%s/%s", keys.BotKey, "sendMessage")
	data := map[string]string{"chat_id": chatID, "text": "hallo frend i got ya maney"}
	jsonData, err := json.Marshal(data)
	if err != nil {
		loggers.ErrorLogger.Println(err)
		return
	}
	//params := url.Values{}
	//params.Add("chat_id", chatID)
	//params.Add("text", "hallo frend i got ya maney")
	//st := params.Encode()
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
	fmt.Println(answer.Status)
	fmt.Println(string(str))
	//TODO:send to admin db payment status from PayoutStatus func like "order_id" , "account_id", "chat_id"(opt, get from связывания таблиц), "status"
	//TODO: create func to return all payments with SUCCESS status
}
