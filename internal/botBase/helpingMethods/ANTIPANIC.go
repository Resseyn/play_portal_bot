package helpingMethods

import "play_portal_bot/pkg/utils/structures"

func IfIsInteracting(chatID int64) bool {
	_, ok := structures.UserStates[chatID]
	return ok
}
