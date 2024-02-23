package auth

import (
	"context"
	"errors"
	database "spectrum300/Database"
	util "spectrum300/Util"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/models"
)

// since https://github.com/vgarvardt/go-oauth2-pg does not do its thing quite as efficiently as i would like,
// i am implementing my own Client Store based off their implementation.
type SpectrumClientStore struct{}

func NewClientStore() (SpectrumClientStore, error) {
	return SpectrumClientStore{}, nil
}

// enter the username to get the client info
func (c SpectrumClientStore) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {

	resultSet, err := database.Pool.Query(
		ctx,
		"SELECT * FROM oauth2_clients WHERE ID = $1",
		id,
	)
	if err != nil {
		return nil, err
	}

	if resultSet.Next() {
		values, err := resultSet.Values()
		if err != nil {
			return nil, err
		}

		return &models.Client{
			ID:     values[0].(string),
			Secret: values[1].(string),
			Domain: values[2].(string),
			Public: true,
			UserID: values[0].(string),
		}, nil
	}

	return nil, errors.New("client not found")
}

func (s *SpectrumClientStore) CreateClient(username string, id int) (oauth2.ClientInfo, error) {

	secret := util.GenerateRandomString(48)

	_, err := database.Pool.Exec(
		context.Background(),
		"INSERT INTO oauth2_clients (ID, Secret, Domain, UserID)\n"+
			"VALUES ($1, $2, $3, $4);",
		username,
		secret,
		"http://localhost", // should this change?
		id,
	)
	if err != nil {
		return &models.Client{}, err
	}

	return &models.Client{
		ID:     username,
		Secret: secret,
		Domain: "http://localhost",
		Public: true,
		UserID: username,
	}, nil
}
