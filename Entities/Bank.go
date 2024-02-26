package entities

import (
	"context"
	"errors"
	"log"
	database "spectrum300/Database"
	"strconv"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
)

// creates a starter bank account for a player's numeric ID
func CreateStarterBankAccountForPlayer(playerID int) error {
	return CreateBankAccountForPlayer(playerID, 1)
}

func CreateBankAccountForPlayer(playerID int, bankID int) error {
	_, err := database.Pool.Exec(
		context.Background(),
		"INSERT INTO Bank_Accounts (BankID, PlayerID, Balance)\n"+
			"VALUES ($1, $2, 0.00);",
		bankID,
		playerID,
	)
	if err != nil {
		log.Println("Failed to create bank account for player with ID:", playerID)
		return err
	}

	return nil
}

type BankAccount struct {
	AccountNumber int             `json:"account_number"`
	BankID        int             `json:"bank_id"`
	Balance       decimal.Decimal `json:"balance"`
}

func GetPlayerBankAccounts(id int) ([]BankAccount, error) {
	rs, err := database.Pool.Query(
		context.Background(),
		"SELECT AccountNumber, BankID, (Balance::NUMERIC)::text\n"+
			"FROM Bank_Accounts\n"+
			"WHERE PlayerID = $1;",
		id,
	)
	if err != nil {
		return nil, err
	}

	var accounts []BankAccount

	for rs.Next() {
		values, err := rs.Values()
		if err != nil {
			return nil, err
		}

		balance, err := decimal.NewFromString(values[2].(string))
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, BankAccount{
			int(values[0].(int32)),
			int(values[1].(int32)),
			balance,
		})
	}

	return accounts, nil
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

// checks if a bank acount belongs to a player
//
// if the query fails, returns false, err
func CheckAccountBelongsToPlayer(id, account int) (bool, error) {
	rs, err := database.Pool.Query(
		context.Background(),
		"SELECT * FROM Bank_Accounts\n"+
			"WHERE PlayerID = $1 AND\n"+
			"AccountNumber = $2;",
		id,
		account,
	)
	if err != nil {
		return false, err
	}

	return rs.Next(), nil
}

func Withdraw(id int, account int, amount decimal.Decimal) error {

	amount_as_string := amount.String()

	num := pgtype.Numeric{}

	err := num.Set(amount_as_string)
	if err != nil {
		log.Println("Failed to convert decimal to pg numeric while withdrawing", err.Error())
		return err
	}

	// actually use transactions here to prevent double writes
	tx, _ := database.Pool.BeginTx(context.Background(), pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	})

	_, err = tx.Exec(
		context.Background(),
		"UPDATE Bank_Accounts\n"+
			"SET Balance = Balance - $1\n"+
			"WHERE AccountNumber = $2;",
		num,
		account,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		context.Background(),
		"UPDATE Players\n"+
			"SET Wallet = Wallet + $1\n"+
			"WHERE ID = $2;",
		num,
		id,
	)
	if err != nil {
		return err
	}

	return tx.Commit(context.Background())
}
