package structures

// UserStates Словарь с взаимодействиями пользователей с ботом (взаимодейсвия, в которых требуется несколько раз что-то ввести и т.д)
// НЕ ИСПОЛЬЗОВАТЬ ДЛЯ ВЗАИМОДЕЙСТВИЯ С ОПЛАТАМИ
var UserStates = make(map[int64]*UserInteraction)

// Moderators is array of chatID's of whom the tickets will be sent
var Moderators = []string{"2038902313", "464644572"}

var UserRedirects = make(map[int64]string)

// Commands - словарь, в котором хранятся коды команд (прим. mainMenu - a1jg; по ключу menu выдаст a1jg)
var Commands = map[string]string{
	"mainMenu":        "aaaa",
	"shop":            "aaab",
	"personalCabinet": "aaac",
	"support":         "aaad",
	"faq":             "aaae",

	"createTicket":    "zxca",
	"respondToTicket": "zxcb", //button for moders only
	"endTicket":       "zxcc",

	"history":    "hist",
	"historyTOP": "hisF",
	"historyBUY": "hisT",

	"topUpBalance":        "aaaf",
	"createCheck":         "aaba",
	"createPayPalychBill": "aabc",

	"respondToOrder": "abaa",
	"endOrder":       "abab",

	"shop_gameServices":       "aaag",
	"shop_services":           "aaah",
	"shop_gameServices_steam": "aaai",

	"spotify":               "aaaj",
	"spotify_individual_1":  "aa1k",
	"spotify_individual_3":  "aa3k",
	"spotify_individual_6":  "aa6k",
	"spotify_individual_12": "a12k",
	"spotify_duo_1":         "duo1",
	"spotify_family_1":      "fam1",

	"spotifySuccessIND1":  "spoa",
	"spotifySuccessIND3":  "spob",
	"spotifySuccessIND6":  "spoc",
	"spotifySuccessIND12": "spod",
	"spotifySuccessDUO1":  "spoe",
	"spotifySuccessFAM1":  "spof",

	"steam_topUpBalance": "aaal",

	"adminPanel":     "aaam",
	"showAdminPanel": "aaao",
	"showReports":    "aaan",

	"pingModer": "ping",
}

var Prices = map[string]float64{
	Commands["spotifySuccessIND1"]:  332.0,
	Commands["spotifySuccessIND3"]:  663.0,
	Commands["spotifySuccessIND6"]:  1243.0,
	Commands["spotifySuccessIND12"]: 1999.0,
	Commands["spotifySuccessDUO1"]:  349.0,
	Commands["spotifySuccessFAM1"]:  402.0,
}

type MessageData struct {
	Command     string
	PrevCommand string
	Price       int
	Custom      string
	//DialogWith int64 //for moders, representing dialog with user
}

type Command struct {
	Text    string
	Command string
}
type UserInteraction struct {
	IsInteracting bool   //optional probably
	Type          string //shop_gameServices_steam_topUpBalance, etc.
	Step          int
	Price         float64
	Order         string
	DataCase      []string
}

//type UserTransaction struct {
//	IsTransacting bool   //optional probably
//	Type          string //shop_gameServices_steam_topUpBalance, etc.
//	Step          int
//	Price         int
//	DataCase      []string
//}
