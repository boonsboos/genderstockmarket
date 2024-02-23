package auth

import (
	"context"
	"errors"
	"log"
	database "spectrum300/Database"
	"time"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/models"
)

// since https://github.com/vgarvardt/go-oauth2-pg does not do its thing quite as efficiently as i would like,
// i am implementing my own Token Store based off their implementation.
type SpectrumTokenStore struct{}

func NewTokenStore() (SpectrumTokenStore, error) {
	return SpectrumTokenStore{}, nil
}

func (store SpectrumTokenStore) Create(ctx context.Context, info oauth2.TokenInfo) error {
	_, err := database.Pool.Exec(
		ctx,
		"INSERT INTO oauth2_tokens (ID, Code, ExpiresAt)\n"+
			"VALUES ($1, $2, $3)",
		info.GetClientID(),
		info.GetCode(),
		info.GetCodeCreateAt().Add(info.GetCodeExpiresIn()),
	)
	if err != nil {
		log.Println("Failed to create token for player with ID=", info.GetClientID())
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

func (store SpectrumTokenStore) RemoveByAccess(ctx context.Context, access string) error {
	return nil
}

func (store SpectrumTokenStore) RemoveByRefresh(ctx context.Context, refresh string) error {
	return nil
}

func (store SpectrumTokenStore) GetByCode(ctx context.Context, code string) (oauth2.TokenInfo, error) {
	rs, err := database.Pool.Query(
		ctx,
		"SELECT * FROM oauth2_tokens WHERE Code = $1",
		code,
	)
	if err != nil {
		log.Println("Failed to get token info by code")
		return nil, err
	}

	if rs.Next() {
		values, err := rs.Values()
		if err != nil {
			log.Println("No records found while searching for token")
			return nil, err
		}

		return &models.Token{
			ClientID:      values[0].(string),
			Code:          values[1].(string),
			CodeCreateAt:  values[2].(time.Time),
			CodeExpiresIn: values[3].(time.Time).Sub(values[2].(time.Time)),
		}, nil
	}

	return nil, errors.New("token not found")
}

func (store SpectrumTokenStore) GetByAccess(ctx context.Context, access string) (oauth2.TokenInfo, error) {
	return &models.Token{}, nil
}

func (store SpectrumTokenStore) GetByRefresh(ctx context.Context, refresh string) (oauth2.TokenInfo, error) {
	return &models.Token{}, nil
}
