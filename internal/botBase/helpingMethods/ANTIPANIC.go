package helpingMethods

import "play_portal_bot/pkg/utils/structures"

func CheckIfIsInteracting(chatID int64) bool {
	_, ok := structures.UserStates[chatID]
	return ok
}
