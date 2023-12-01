package databaseModels

import (
	"database/sql"
)

var NotEnoghtCash error

type UsersDB struct {
	DB *sql.DB
}

var Users UsersDB

type UserInfo struct {
	ChatID  int64
	Balance float64
}

func (m *UsersDB) GetUser(chatID int64) (*UserInfo, error) {
	var balance float64
	err := m.DB.QueryRow("SELECT balance FROM user_info WHERE chat_id = $1", chatID).Scan(&balance)
	if err != nil {
		return nil, err
	}
	found := &UserInfo{
		ChatID:  chatID,
		Balance: balance,
	}
	return found, nil
}

func (m *UsersDB) CreateUser(chatID int64) (*UserInfo, error) {
	newUser := &UserInfo{
		ChatID:  chatID,
		Balance: 0,
	}
	_, err := m.DB.Exec("INSERT INTO user_info (chat_id, balance) VALUES ($1, $2)",
		newUser.ChatID, newUser.Balance)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (m *UsersDB) TopUpBalance(chatID int64, amount float64) (float64, error) {
	user, _ := m.GetUser(chatID)
	_, err := m.DB.Exec("UPDATE user_info SET balance = balance + $2 WHERE chat_id = $1",
		chatID, amount)
	if err != nil {
		return 0, err
	}

	return user.Balance + amount, nil
}

func (m *UsersDB) ConsumeBalance(chatID int64, amount float64) (float64, error) {
	user, _ := m.GetUser(chatID)
	balance := user.Balance
	if balance-amount >= 0 {
		_, err := m.DB.Exec("UPDATE user_info SET balance = balance - $2 WHERE chat_id = $1",
			chatID, amount)
		if err != nil {
			return 0, err
		}

		return balance - amount, nil
	} else {
		return user.Balance, NotEnoghtCash
	}

}
