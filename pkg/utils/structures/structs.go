package structures

// UserStates Словарь с взаимодействиями пользователей с ботом (взаимодейсвия, в которых требуется несколько раз что-то ввести и т.д)
var UserStates = make(map[int64]*UserInteraction)

type MessageData struct {
	ChatID      int64
	MessageID   int
	Command     string
	PrevCommand string
}

type Command struct {
	Text    string
	Command string
}
type UserInteraction struct {
	IsInteracting bool   //optional probably
	Type          string //shop_gameServices_steam_topUpBalance, etc.
	Step          int
	DataCase      []string
}
