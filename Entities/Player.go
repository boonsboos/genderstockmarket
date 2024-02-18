package entities

// contains minimal user information
type PlayerInfo struct {
	Name     string `json:"name"`
	NetWorth uint64 `json:"net_worth"` // should be fine for now
}

// contains the full data about the player
type Player struct {
	// ID should not be shown to the API user,
	// just for internal purposes
	ID uint64 `json:"userID"`
	PlayerInfo
}
