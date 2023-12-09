package databaseModels

import (
	"database/sql"
	"fmt"
	"play_portal_bot/pkg/utils/structures"
	"time"
)

type OrdersDB struct {
	DB *sql.DB
}

var Orders OrdersDB

type DBOrder struct {
	ChatID  int64
	OrderID string
	Amount  float64
	Custom  string
	Status  string
	Data    string
}

func (m *OrdersDB) GetOrder(orderID string) (*DBOrder, error) {
	found := &DBOrder{}
	err := m.DB.QueryRow("SELECT * FROM orders WHERE order_id = $1", orderID).Scan(&found.ChatID, &found.OrderID, &found.Amount, &found.Custom, &found.Status, &found.Data)
	if err != nil {
		return nil, err
	}
	return found, nil
}

func (m *OrdersDB) CreateOrder(chatID int64, orderID string, amount float64, custom string) (*DBOrder, error) {
	newOrder := &DBOrder{
		ChatID:  chatID,
		OrderID: orderID,
		Amount:  amount,
		Custom:  custom,
		Status:  "NEW",
		Data:    time.Now().Format("02.01.2006 15:04"),
	}
	_, err := m.DB.Exec("INSERT INTO orders (chat_id, order_id, amount, custom, status, data) VALUES ($1, $2, $3, $4, $5, $6)",
		newOrder.ChatID, newOrder.OrderID, newOrder.Amount, newOrder.Custom, newOrder.Status, newOrder.Data)
	if err != nil {
		return nil, err
	}
	return newOrder, nil
}

// OrderIsDone creates check in DB
func (m *OrdersDB) OrderIsDone(orderID string) (*DBOrder, error) {
	order, _ := m.GetOrder(orderID)
	_, err := m.DB.Exec("UPDATE orders SET status = $1 WHERE order_id = $2", "SUCCESS", orderID)
	if err != nil {
		return nil, err
	}
	_, err = m.DB.Exec("INSERT INTO checks (chat_id, order_id, amount, custom, data) VALUES ($1, $2, $3, $4, $5)",
		order.ChatID, order.OrderID, order.Amount, order.Custom, order.Data)
	if err != nil {
		return nil, err
	}
	return nil, nil //TODO: ну эт хуйня хд
}

//// CreateCheck
//func (m *OrdersDB) CreateCheck(orderID string, chatID int64, amount float64, custom string) (*DBOrder, error) {
//	data := time.Now().Format("02.01.2006 15:04")
//	_, err := m.DB.Exec("INSERT INTO checks (chat_id, order_id, amount, custom, data) VALUES ($1, $2, $3, $4, $5)",
//		chatID, orderID, amount, custom, data)
//	if err != nil {
//		return nil, err
//	}
//	return nil, nil //TODO: ну эт хуйня хд
//}

func (m *OrdersDB) DeletePrevOrderIfPresent(chatID int64) {
	found := &DBOrder{}
	m.DB.QueryRow("SELECT * FROM orders WHERE chat_id = $1", chatID).Scan(&found.ChatID, &found.OrderID, &found.Amount, &found.Custom, &found.Status, &found.Data)
	if found.Status == "NEW" && found.Custom == "aaac" {
		m.DB.Exec("DELETE FROM orders WHERE chat_id = $1 AND custom = 'aaac' AND status = 'NEW'", chatID)
	}
}

// ShowOrdersHistory выводит пользователю историю пополнений и покупок. false, если вывести историю пополнений, true, если историю покупок
func (m *OrdersDB) ShowOrdersHistory(chatID int64, orderHistory bool) string {
	var amount float64
	var orders string
	if !orderHistory {
		orders = "Ваша история пополнений:\n\n"
		rows, _ := m.DB.Query("SELECT (amount) FROM checks WHERE chat_id = $1 AND custom = 'aaac' ORDER BY data DESC LIMIT 10", chatID)
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&amount)
			orders += fmt.Sprintf("Пополнение на %v рублей\n", amount)
		}
		if orders == "Ваша история пополнений:\n\n" {
			return "У вас еще не было пополнений!"
		}
	} else {
		orders = "Ваша история покупок:\n\n"
		var custom string
		rows, _ := m.DB.Query("SELECT custom, amount FROM checks WHERE chat_id = $1 AND custom <> 'aaac' ORDER BY data DESC LIMIT 10", chatID)
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&custom, &amount)
			orders += fmt.Sprintf("Покупка %v за %v рублей\n", structures.Codes[custom], amount)
		}
		if orders == "Ваша история покупок:\n\n" {
			return "У вас еще не было покупок!"
		}
	}
	return orders
}
