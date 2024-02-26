package auth

import (
	"context"
	"errors"
	"log"
	database "spectrum300/Database"
	"time"
)

// since https://github.com/vgarvardt/go-oauth2-pg does not do its thing quite as efficiently as i would like,
// i am implementing my own Token Store based off their implementation.
type SpectrumTokenStore struct{}

func NewSpectrumTokenStore() (SpectrumTokenStore, error) {
	return SpectrumTokenStore{}, nil
}

func (store SpectrumTokenStore) Create(ctx context.Context, token Token) error {
	_, err := database.Pool.Exec(
		ctx,
		"INSERT INTO oauth2_tokens (ID, Code, ExpiresAt)\n"+
			"VALUES ($1, $2, $3)\n"+
			"ON CONFLICT (ID)\n"+
			"DO UPDATE SET Code = EXCLUDED.Code,"+
			"CreatedAt = EXCLUDED.CreatedAt,"+
			"ExpiresAt = EXCLUDED.ExpiresAt WHERE oauth2_tokens.ID = $1;", // if token gets re-requested before expiry
		token.ID,
		token.Code,
		token.ExpiresAt,
	)
	if err != nil {
		log.Println("Failed to create token for player with ID:", token.ID)
		log.Println(err.Error())
		return err
	}

	return nil
}

func (store SpectrumTokenStore) RemoveByCode(ctx context.Context, code string) error {
	_, err := database.Pool.Exec(
		ctx,
		"DELETE FROM oauth2_tokens WHERE Code = $1",
		code,
	)
	if err != nil {
		log.Println("Failed to delete token by code")
		return err
	}

	return nil
}

func (store SpectrumTokenStore) GetByCode(ctx context.Context, code string) (Token, error) {
	rs, err := database.Pool.Query(
		ctx,
		"SELECT * FROM oauth2_tokens WHERE Code = $1",
		code,
	)
	if err != nil {
		log.Println("Failed to get token info by code")
		return Token{}, err
	}

	if rs.Next() {
		values, err := rs.Values()
		if err != nil {
			log.Println("No records found while searching for token")
			return Token{}, err
		}

		return Token{
			ID:        int(values[0].(int32)),
			Code:      values[1].(string),
			CreatedAt: values[2].(time.Time),
			ExpiresAt: values[3].(time.Time),
		}, nil
	}

	return Token{}, errors.New("token not found")
}
