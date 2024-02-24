package auth

import (
	"context"
	"errors"
	database "spectrum300/Database"
)

// since https://github.com/vgarvardt/go-oauth2-pg does not do its thing quite as efficiently as i would like,
// i am implementing my own Client Store based off their implementation.
type SpectrumClientStore struct{}

func NewSpectrumClientStore() (SpectrumClientStore, error) {
	return SpectrumClientStore{}, nil
}

// enter the username to get the client info
func (c SpectrumClientStore) GetByID(ctx context.Context, id string) (Client, error) {

	resultSet, err := database.Pool.Query(
		ctx,
		"SELECT * FROM oauth2_clients WHERE ID = $1",
		id,
	)
	if err != nil {
		return Client{}, err
	}

	if resultSet.Next() {
		values, err := resultSet.Values()
		if err != nil {
			return Client{}, err
		}

		return Client{
			ID:     values[0].(string),
			Secret: values[1].(string),
			Domain: values[2].(string),
			UserID: int(values[3].(int32)),
		}, nil
	}

	return Client{}, errors.New("client not found")
}

func (s *SpectrumClientStore) CreateClient(username string, id int) (Client, error) {

	secret := GenerateSecret()

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
		return Client{}, err
	}

	return Client{
		ID:     username,
		Secret: secret,
		Domain: "http://localhost",
		UserID: id,
	}, nil
}
