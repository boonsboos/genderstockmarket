package entities

import (
	"context"
	"errors"
	"log"
	auth "spectrum300/Auth"
	database "spectrum300/Database"
	"strconv"

	"github.com/shopspring/decimal"
)

// contains minimal user information
type PlayerInfo struct {
	Name     string          `json:"name"`
	NetWorth decimal.Decimal `json:"net_worth"` // should be fine for now
}

// contains the full data about the player
type Player struct {
	// ID should not be shown to the API user,
	// just for internal purposes
	ID uint64 `json:"userID"`
	PlayerInfo
}

// saves a new player to the database, creating an oauth client at the same time
func SaveNewPlayer(name string) error {

	_, err := auth.ClientStore.CreateClient(name)
	if err != nil {
		log.Println("Failed to save new OAuth client:", err.Error())
		return err
	}

	_, err = database.DatabaseConnection.Exec(context.Background(),
		"INSERT INTO Players (Username, NetWorth)\n"+
			"VALUES ($1, 100.00);",
		name,
	)
	if err != nil {
		log.Println("Failed to save new user:", err.Error())
		return err
	}

	return nil
}

// returns the player's total balance across all their bank accounts
func GetPlayerTotalBalance(id int) (decimal.Decimal, error) {
	rs, err := database.DatabaseConnection.Query(
		context.Background(),
		"SELECT SUM(Balance::NUMERIC)::text AS Total FROM Bank_Accounts WHERE PlayerID = $1;",
		id,
	)
	if err != nil {
		log.Println("Failed to get total balance for player with ID=", id)
		return decimal.Zero, err
	}

	if rs.Next() {
		values, err := rs.Values()
		if err != nil {
			log.Println("Failed to get values while getting total balance for player with ID=", id)
			return decimal.Zero, err
		}

		return decimal.NewFromString(values[0].(string))
	}

	return decimal.Zero, errors.New("player with ID=" + strconv.Itoa(id) + " not found")
}
