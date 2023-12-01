package databaseModels

import (
	"database/sql"
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
}

func (m *OrdersDB) GetOrder(orderID string) (*DBOrder, error) {
	found := &DBOrder{}
	err := m.DB.QueryRow("SELECT * FROM orders WHERE order_id = $1", orderID).Scan(&found.ChatID, &found.OrderID, &found.Amount, &found.Custom, &found.Status)
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
	}
	_, err := m.DB.Exec("INSERT INTO orders (chat_id, order_id, amount, custom, status) VALUES ($1, $2, $3, $4, $5)",
		newOrder.ChatID, newOrder.OrderID, newOrder.Amount, newOrder.Custom, newOrder.Status)
	if err != nil {
		return nil, err
	}
	return newOrder, nil
}
