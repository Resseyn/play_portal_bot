package databaseModels

import (
	"database/sql"
	"fmt"
	"play_portal_bot/pkg/utils/structures"
)

type KeysDB struct {
	DB *sql.DB
}

var Keys KeysDB

type DBkey struct {
	Key       string
	KeyCode   string
	Avaliable bool
}

func (m *KeysDB) GetKey(orderCode string) (*DBkey, error) {
	structures.GlobalMutex.Lock()
	found := &DBkey{}
	err := m.DB.QueryRow("SELECT * FROM keys WHERE key_code = $1", orderCode).Scan(&found.Key, &found.KeyCode, &found.Avaliable)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	_, err = m.DB.Exec("DELETE FROM keys WHERE key = $1", found.Key)
	if err != nil {
		return nil, err
	}
	structures.GlobalMutex.Unlock()
	fmt.Println(found)
	return found, nil
}
