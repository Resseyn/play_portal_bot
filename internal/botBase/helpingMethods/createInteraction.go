package helpingMethods

import "play_portal_bot/pkg/utils/structures"

// NewInteraction creates new interaciton for user, optPrice and optData is optional
func NewInteraction(interactionType string, chatID int64, optPrice float64, optData []string) {
	delete(structures.UserStates, chatID)
	structures.UserStates[chatID] = &structures.UserInteraction{
		IsInteracting: true,
		Type:          interactionType,
		Step:          0,
		Price:         optPrice,
		DataCase:      optData,
	}
}
