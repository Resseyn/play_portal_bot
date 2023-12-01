package play_portal_bot

import (
	"play_portal_bot/internal/databaseModels"
	"play_portal_bot/internal/loggers"
	"play_portal_bot/pkg/database"
	"testing"
)

func Test(t *testing.T) {
	loggers.InitLogger()
	database.InitDatabase()
	databaseModels.InitModels()
	_, err := databaseModels.Users.CreateUser(298)
	if err != nil {
		t.Error(err)
	}
	user, _ := databaseModels.Users.GetUser(298)
	t.Log(*user)
	balance, _ := databaseModels.Users.TopUpBalance(298, 100)
	t.Log(balance)
	balance, _ = databaseModels.Users.ConsumeBalance(298, 70)
	t.Log(balance)
	balance, err = databaseModels.Users.ConsumeBalance(298, 70)
	if err != nil {
		t.Error(err)
	}
	t.Log(balance)
	_, err = databaseModels.Orders.CreateOrder(298, "XUY2287714881488", 200, "aabb")
	if err != nil {
		t.Error(err)
	}
	_, err = databaseModels.Orders.CreateOrder(297, "XUY2287714881488", 200, "aabb")
	if err != nil {
		t.Error(err)
	}
	_, err = databaseModels.Orders.CreateOrder(298, "XUY2287714881487", 200, "aabb")
	if err != nil {
		t.Error(err)
	}
	order, err := databaseModels.Orders.GetOrder("XUY2287714881488")
	if err != nil {
		t.Error(err)
	}
	t.Log(*order)
}
