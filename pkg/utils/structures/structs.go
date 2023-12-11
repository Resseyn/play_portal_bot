package structures

import (
	"sync"
)

// UserStates Словарь с взаимодействиями пользователей с ботом (взаимодейсвия, в которых требуется несколько раз что-то ввести и т.д)
// НЕ ИСПОЛЬЗОВАТЬ ДЛЯ ВЗАИМОДЕЙСТВИЯ С ОПЛАТАМИ
var GlobalMutex sync.Mutex

var UserStates = make(map[int64]*UserInteraction)

// Goods - key - название товара, value - 0 - главное меню товара [0] - текст, [1] - url картинки, 1 - первый товар [0] -
// название кнопки, [1] - код товара, [2] - url на картинку, [3] - текст внутри товара

type TypicalPage struct {
	URL      string
	Text     string
	Commands [][]Command
	Data     *MessageData
	PrevPage string
	Custom   string
}

//nodeMap - node, {connectedNodes}   	nodeInfo - node - params(столбики всякие) ({connectedNodes})

// Moderators are array of chatID's of whom the tickets will be sent
var Moderators = []string{"2038902313", "464644572"}

// UserOrders - Словарь для редиректа у кнопки, появляющейся после пополнения счета. содержит код заказа
var UserOrders = make(map[int64]string)

// Handlers - словарь, который коду товара присуждает хэндлер для обработки его покупки
var Handlers = map[string]string{

	"spoa": "spotifyHandler",
	"spob": "spotifyHandler",
	"spoc": "spotifyHandler",
	"spod": "spotifyHandler",
	"spoe": "spotifyHandler",
	"spof": "spotifyHandler",

	"app1": "keyHandler",
	"app2": "keyHandler",
	"app3": "keyHandler",
	"app4": "keyHandler",
	"app5": "keyHandler",
}

var Parameters = map[string][]string{

	"spotifyHandler": {"Введите логин от Spotify", "Введите пароль от Spotify", "cкинь писю"},
	"keyHandler":     {},
}

var Pages = map[string]*TypicalPage{
	Commands["shop"]: &TypicalPage{
		URL:  "pkg/utils/data/img/mainMenuImages/gettyimages-1067956982.jpg",
		Text: "МАГАЗИН ИГР 'ЗМЕЙ ГЕРОИНЫЧ' через пэйдеж",
		Commands: [][]Command{
			{
				{Text: "Игровые сервисы", Command: Commands["shop_gameServices"]},
				{Text: "Сервисы", Command: Commands["shop_services"]}},
			{
				{Text: "Pepega(насвай не завезли)", Command: Commands[""]},
			}},
		PrevPage: Commands["mainMenu"],
	},
	Commands["shop_gameServices"]: &TypicalPage{
		URL:  "pkg/utils/data/img/shopImages/gameServices.jpg",
		Text: "Выберите категорию",
		Commands: [][]Command{
			{
				{Text: "Steam", Command: ""},
				{Text: "Xbox/Microsoft", Command: ""}},
			{
				{Text: "Playstation", Command: ""},
				{Text: "ИГРЫ", Command: ""}},
		},
		PrevPage: Commands["shop"],
	},
	Commands["shop_services"]: &TypicalPage{
		URL:  "pkg/utils/data/img/shopImages/gameServices.jpg",
		Text: "Выберите категорию",
		Commands: [][]Command{
			{
				//{Text: "Подписка Twitch", Command: ""},
				{Text: "Spotify", Command: Commands["spotify"]}},
			{
				{Text: "AppStore", Command: Commands["appStore"]}},
		},
		PrevPage: Commands["shop"],
	},
	//========SPOTIFY==========//========SPOTIFY==========//========SPOTIFY==========
	Commands["spotify"]: &TypicalPage{
		URL:  "pkg/utils/data/img/shopImages/servicesImages/spotify.jpg",
		Text: "Выберите товар:",
		Commands: [][]Command{
			{
				{Text: "Spotify Individual 1 месяц", Command: Commands["spotify_individual_1"]}},
			{
				{Text: "Spotify Individual 3 месяца", Command: Commands["spotify_individual_3"]}},
			{
				{Text: "Spotify Individual 6 месяцев", Command: Commands["spotify_individual_6"]}},
			{
				{Text: "Spotify Individual 12 месяцев", Command: Commands["spotify_individual_12"]}},
			{
				{Text: "Spotify DUO 1 месяц", Command: Commands["spotify_duo_1"]}},
			{
				{Text: "Spotify Family 1 месяц", Command: Commands["spotify_family_1"]}},
		},
		PrevPage: Commands["shop_services"],
	},
	//TODO: может можно вариации товара как-то объединить???
	Commands["spotify_individual_1"]: &TypicalPage{
		URL:  "pkg/utils/data/img/shopImages/servicesImages/spotify/spotify_individual_1.jpg",
		Text: "Выберите товар", //TODO: описание добавить
		Commands: [][]Command{
			{
				{Text: "Купить", Command: Commands["topUpBalance"]}},
		},
		PrevPage: Commands["spotify"],
		Custom:   "spoa",
	},
	Commands["spotify_individual_3"]: &TypicalPage{
		URL:  "pkg/utils/data/img/shopImages/servicesImages/spotify/spotify_individual_3.jpg",
		Text: "Выберите товар", //TODO: описание добавить
		Commands: [][]Command{
			{
				{Text: "Купить", Command: Commands["topUpBalance"]}},
		},
		PrevPage: Commands["spotify"],
		Custom:   "spob",
	},
	Commands["spotify_individual_6"]: &TypicalPage{
		URL:  "pkg/utils/data/img/shopImages/servicesImages/spotify/spotify_individual_6.jpg",
		Text: "Выберите товар", //TODO: описание добавить
		Commands: [][]Command{
			{
				{Text: "Купить", Command: Commands["topUpBalance"]}},
		},
		PrevPage: Commands["spotify"],
		Custom:   "spoc",
	},
	Commands["spotify_individual_12"]: &TypicalPage{
		URL:  "pkg/utils/data/img/shopImages/servicesImages/spotify/spotify_individual_12.jpg",
		Text: "Выберите товар", //TODO: описание добавить
		Commands: [][]Command{
			{
				{Text: "Купить", Command: Commands["topUpBalance"]}},
		},
		PrevPage: Commands["spotify"],
		Custom:   "spod",
	},
	Commands["spotify_duo_1"]: &TypicalPage{
		URL:  "pkg/utils/data/img/shopImages/servicesImages/spotify/spotify_DUO_1.jpg",
		Text: "Выберите товар", //TODO: описание добавить
		Commands: [][]Command{
			{
				{Text: "Купить", Command: Commands["topUpBalance"]}},
		},
		PrevPage: Commands["spotify"],
		Custom:   "spoe",
	},
	Commands["spotify_family_1"]: &TypicalPage{
		URL:  "pkg/utils/data/img/shopImages/servicesImages/spotify/spotify_family_1.jpg",
		Text: "Выберите товар", //TODO: описание добавить
		Commands: [][]Command{
			{
				{Text: "Купить", Command: Commands["topUpBalance"]}},
		},
		PrevPage: Commands["spotify"],
		Custom:   "spof",
	},
	//========SPOTIFY==========//========SPOTIFY==========//========SPOTIFY==========

	Commands["faq"]: &TypicalPage{
		URL:      "pkg/utils/data/img/mainMenuImages/faq.png",
		Text:     "Здесь можно почитать ответы на Часто задаваемые вопросы. НУ И ТИПО ССЫЛОЧКУ СЮДА АГА",
		Commands: [][]Command{{}},
		PrevPage: Commands["mainMenu"],
	},
	Commands["history"]: &TypicalPage{
		URL:  "pkg/utils/data/img/mainMenuImages/faq.png",
		Text: "Какую историю желаете посмотреть?",
		Commands: [][]Command{
			{
				{Text: "Историю пополнений", Command: Commands["historyTOP"]}},
			{
				{Text: "Историю покупок", Command: Commands["historyBUY"]},
			}},
		PrevPage: Commands["personalCabinet"],
	},
}

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

	"Success": "spot",
	// я изменял хуйню с миллиономами хэндлеров, теперь вся инфа - это код заказа, которой декодируется в название, и его цена
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
	"spoe": 354.0,
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
