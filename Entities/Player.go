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
	_, err := database.Pool.Exec(context.Background(),
		"INSERT INTO Players (Username, NetWorth)\n"+
			"VALUES ($1, 100.00) ON CONFLICT DO NOTHING;",
		name,
	)
	if err != nil {
		log.Println("Failed to save new user:", err.Error())
		return err
	}

	// get the player ID
	id, err := GetPlayerIDByUsername(name)
	if err != nil {
		log.Println("Failed to save new user, ID not found? ", err.Error())
		return err
	}

	log.Println("saving new oauth client")
	_, err = auth.ClientStore.CreateClient(name, id)
	if err != nil {
		log.Println("Failed to save new OAuth client:", err.Error())
		return err
	}

	return nil
}

func GetPlayerIDByUsername(name string) (int, error) {
	rs, err := database.Pool.Query(
		context.Background(),
		"SELECT ID FROM Players WHERE Username = $1",
		name,
	)
	if err != nil {
		log.Println("Failed to query ID for player with Username=", name)
		return -1, errors.New("Player not found")
	}
	if rs.Next() {
		values, err := rs.Values()
		if err != nil {
			log.Println("Failed to get values while getting ID of player with Username=", name)
			return -1, err
		}

		return int(values[0].(int32)), nil
	}

	return -1, errors.New("Player not found")
}

// returns the player's total balance across all their bank accounts
func GetPlayerTotalBalance(id int) (decimal.Decimal, error) {
	rs, err := database.Pool.Query(
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
