package structures

// UserStates Словарь с взаимодействиями пользователей с ботом (взаимодейсвия, в которых требуется несколько раз что-то ввести и т.д)
// НЕ ИСПОЛЬЗОВАТЬ ДЛЯ ВЗАИМОДЕЙСТВИЯ С ОПЛАТАМИ
var UserStates = make(map[int64]*UserInteraction)

// Moderators is array of chatID's of whom the tickets will be sent
var Moderators = []string{"2038902313", "464644572"}

// UserRedirectsAndOrders - Словарь для редиректа у кнопки, появляющейся после пополнения счета. содержит нужный хендлер [0]
// и собстевнно код заказа [1]
var UserRedirectsAndOrders = make(map[int64][]string)

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

	"appStore":     "apps",
	"appStore500":  "appa",
	"appStore1000": "appb",
	"appStore1500": "appc",
	"appStore3000": "appd",
	"appStore9000": "appe",

	"spotifySuccess": "spot",
	//TODO: я изменял хуйню с миллиономами хэндлеров, теперь вся инфа - это код заказа, которой декодируется в название, и его цена
	"steam_topUpBalance": "aaal",

	"adminPanel":     "aaam",
	"showAdminPanel": "aaao",
	"showReports":    "aaan",

	"pingModer": "ping",
}

var Prices = map[string]float64{

	"spoa": 332.0,
	"spob": 663.0,
	"spoc": 1243.0,
	"spod": 2017.0,
	"spoe": 349.0,
	"spof": 402.0,

	"app1": 689.0,
	"app2": 1379.0,
	"app3": 2068.0,
	"app4": 4136.0,
	"app5": 12408.0,
}

var Codes = map[string]string{

	"spoa": "Spotify Individual 1 месяц",
	"spob": "Spotify Individual 3 месяца",
	"spoc": "SpotifyIndividual6",
	"spod": "SpotifyIndividual12",
	"spoe": "SpotifyDUO1",
	"spof": "SpotifyFAMILY1",

	"app1": "Ключ AppStore 500 руб",
	"app2": "Ключ AppStore 1000 руб",
	"app3": "Ключ AppStore 1500 руб",
	"app4": "Ключ AppStore 3000 руб",
	"app5": "Ключ AppStore 9000 руб",
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
