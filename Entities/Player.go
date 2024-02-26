package entities

import (
	"context"
	"errors"
	"log"
	auth "spectrum300/Auth"
	database "spectrum300/Database"

	"github.com/jackc/pgtype"
	"github.com/shopspring/decimal"
)

// contains the full data of the player entity
type Player struct {
	// ID should not be shown to the API user,
	// just for internal purposes
	ID     int             `json:"userID"`
	Name   string          `json:"username"`
	Wallet decimal.Decimal `json:"wallet"`
}

// saves a new player to the database, creating an oauth client at the same time
func SaveNewPlayer(name string) error {
	_, err := database.Pool.Exec(context.Background(),
		"INSERT INTO Players (Username)\n"+
			"VALUES ($1) ON CONFLICT DO NOTHING;",
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

	_, err = auth.ClientStore.CreateClient(name, id)
	if err != nil {
		log.Println("Failed to save new OAuth client:", err.Error())
		return err
	}

	err = CreateStarterBankAccountForPlayer(id)
	if err != nil {
		log.Println("Failed to save starter bank account for player with ID:", id, err.Error())
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

type PlayerStats struct {
	Username           string          `json:"username"`
	TotalTradesDone    uint64          `json:"totalTrades"`
	TotalTradingVolume decimal.Decimal `json:"totalTradingVolume"`
	Firm               string
}

func GetPlayerProfile(id int) (Player, error) {
	rs, err := database.Pool.Query(
		context.Background(),
		"SELECT ID, Username, Wallet FROM Players WHERE ID = $1",
		id,
	)
	if err != nil {
		log.Println("Failed to get profile for Player:", id)
		return Player{}, errors.New("Player not found")
	}
	if rs.Next() {
		values, err := rs.Values()
		if err != nil {
			return Player{}, err
		}

		var wallet_str string
		n := values[2].(pgtype.Numeric)
		err = n.AssignTo(&wallet_str)
		if err != nil {
			log.Println("Failed to convert db numeric to string")
			return Player{}, err
		}

		wallet, err := decimal.NewFromString(wallet_str)
		if err != nil {
			log.Println("kanker error", err.Error())
			return Player{}, err
		}

		return Player{
			int(values[0].(int32)),
			values[1].(string),
			wallet,
		}, nil
	}

	return Player{}, errors.New("Player not found")
}
