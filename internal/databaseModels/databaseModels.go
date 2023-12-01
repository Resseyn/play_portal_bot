package databaseModels

import "play_portal_bot/pkg/database"

func InitModels() {
	Users.DB = database.GlobalDatabase
	Orders.DB = database.GlobalDatabase
}
