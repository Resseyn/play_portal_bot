package helpingMethods

import (
	"gopkg.in/telebot.v3"
)

func CreateOrder(c telebot.Context) error {
	c.Send("order created")
	return nil
}
