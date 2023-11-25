package structures

// UserStates Словарь с взаимодействиями пользователей с ботом (взаимодейсвия, в которых требуется несколько раз что-то ввести и т.д)
var UserStates = make(map[int64]*UserInteraction)

// Commands - словарь, в котором хранятся коды команд (прим. mainMenu - a1jg; по ключу menu выдаст a1jg)
var Commands = map[string]string{
	"mainMenu":                "a1jg",
	"shop":                    "n4kj",
	"personalCabinet":         "g3p3",
	"support":                 "13nk",
	"faq":                     "fkr1",
	"buy":                     "nfa2",
	"shop_gameServices":       "1nd3",
	"shop_services":           "1ar4",
	"shop_gameServices_steam": "d1z7",
	"spotify":                 "13l1",
	"spotify_individual_1":    "fb41",
	"steam_topUpBalance":      "1bof",
}

type MessageData struct {
	ChatID      int64
	MessageID   int
	Command     string
	PrevCommand string
	Price       int
}

type Command struct {
	Text    string
	Command string
}
type UserInteraction struct {
	IsInteracting bool   //optional probably
	Type          string //shop_gameServices_steam_topUpBalance, etc.
	Step          int
	Price         int
	DataCase      []string
}

//type UserTransaction struct {
//	IsTransacting bool   //optional probably
//	Type          string //shop_gameServices_steam_topUpBalance, etc.
//	Step          int
//	Price         int
//	DataCase      []string
//}
