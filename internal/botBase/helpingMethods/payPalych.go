package helpingMethods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/telebot.v3"
	"io/ioutil"
	"log"
	"net/http"
	"play_portal_bot/internal/loggers"
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
