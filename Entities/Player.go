package entities

import (
	"context"
	"log"
	auth "spectrum300/Auth"
	database "spectrum300/Database"
	util "spectrum300/Util"
)

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

// saves a new player to the database, creating an oauth client at the same time
func SaveNewPlayer(name string) {
	auth.ClientStore.Create(client{
		ID:     name, // TODO: change this to uniform size hash like sha512
		Secret: util.GenerateRandomString(32),
		Domain: "http://localhost",
		Data:   make([]byte, 0),
	})

	_, err := database.DatabaseConnection.Conn.Query(context.Background(),
		"INSERT INTO Players (Username, NetWorth)\n"+
			"VALUES ($1, 100.00);",
		name,
	)
	if err != nil {
		log.Println("Failed to save new user:", err.Error())
	}
}
