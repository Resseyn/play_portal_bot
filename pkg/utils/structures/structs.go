package structures

// UserStates Словарь с взаимодействиями пользователей с ботом (взаимодейсвия, в которых требуется несколько раз что-то ввести и т.д)
var UserStates = make(map[int64]*UserInteraction)

// Commands - словарь, в котором хранятся коды команд (прим. mainMenu - a1jg; по ключу menu выдаст a1jg)
var Commands = map[string]string{
	"mainMenu":                "aaaa",
	"shop":                    "aaab",
	"personalCabinet":         "aaac",
	"support":                 "aaad",
	"faq":                     "aaae",
	"buy":                     "aaaf",
	"shop_gameServices":       "aaag",
	"shop_services":           "aaah",
	"shop_gameServices_steam": "aaai",
	"spotify":                 "aaaj",
	"spotify_individual_1":    "aaak",
	"steam_topUpBalance":      "aaal",
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
