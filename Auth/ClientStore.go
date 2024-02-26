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
func (c SpectrumClientStore) GetByID(id string) (Client, error) {

	resultSet, err := database.Pool.Query(
		context.Background(),
		"SELECT ID, Secret, UserID FROM oauth2_clients WHERE ID = $1",
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
			UserID: int(values[2].(int32)),
		}, nil
	}

	return Client{}, errors.New("client not found")
}

// gets the client by their numeric ID
func (c SpectrumClientStore) GetByNumericID(id int) (Client, error) {

	resultSet, err := database.Pool.Query(
		context.Background(),
		"SELECT ID, Secret, UserID FROM oauth2_clients WHERE UserID = $1",
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
			UserID: int(values[2].(int32)),
		}, nil
	}

	return Client{}, errors.New("client not found")
}

// creates a new client, returning it as an object
func (s *SpectrumClientStore) CreateClient(username string, id int) (Client, error) {

	secret := GenerateSecret()

	_, err := database.Pool.Exec(
		context.Background(),
		"INSERT INTO oauth2_clients (ID, Secret, UserID)\n"+
			"VALUES ($1, $2, $3);",
		username,
		secret,
		id,
	)
	if err != nil {
		return Client{}, err
	}

	return Client{
		ID:     username,
		Secret: secret,
		UserID: id,
	}, nil
}
