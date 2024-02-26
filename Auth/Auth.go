package auth

import (
	"context"
	"errors"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var TokenStore SpectrumTokenStore
var ClientStore SpectrumClientStore

var _aa interface{} // do not remove

func InitAuthServer(router *gin.Engine) {
	TokenStore, err := NewSpectrumTokenStore()
	if err != nil {
		log.Fatal("Failed to created token store:", err.Error())
	}
	log.Println("Token store OK")

	ClientStore, err = NewSpectrumClientStore()
	if err != nil {
		log.Fatal("Failed to created client store:", err.Error())
	}
	log.Println("Client store OK")

	// go compiler cries otherwise, no clue why
	_aa = struct {
		t SpectrumTokenStore
	}{
		TokenStore,
	}

	log.Println("Auth server OK")

	// token route is getting registered
}

type Token struct {
	ID        int
	Code      string
	CreatedAt time.Time
	ExpiresAt time.Time
}

type Client struct {
	ID     string
	Secret string
	UserID int
}

const prefix = "Bearer "

func ValidateBearerToken(header string) (Token, error) {
	if !strings.HasPrefix(header, prefix) {
		return Token{}, MalformattedAuthorizationHeader
	}

	code := strings.Split(header, " ")

	return TokenStore.GetByCode(context.Background(), code[1])
}

func ValidateTokenRequest(params url.Values) (string, string, error) {
	client := params.Get("client_id")
	if client == "" {
		return "", "", errors.New("client_id not found")
	}

	secret := params.Get("client_secret")
	if secret == "" {
		return "", "", errors.New("client_secret not found")
	}

	grantType := params.Get("grant_type")
	if grantType != "client_credentials" {
		return "", "", errors.New("this grant type is not supported")
	}

	clientInfo, err := ClientStore.GetByID(client)
	if err != nil {
		log.Println("Failed to get client in Token Request:", err.Error())
		return "", "", errors.New("internal server error")
	}

	if clientInfo.ID == client && clientInfo.Secret == secret {
		return client, secret, nil
	}

	return "", "", errors.New("not authorized")
}

func CreateNewToken(client, secret string) (Token, error) {

	clientInfo, err := ClientStore.GetByID(client)
	if err != nil {
		log.Println("Failed to get ID of player while creating new token")
		return Token{}, err
	}

	token := Token{
		ID:        clientInfo.UserID,
		Code:      GenerateToken(client, secret),
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Hour * 4),
	}

	// insert into db
	err = TokenStore.Create(context.Background(), token)
	if err != nil {
		log.Println("Failed to save token for ID:", client)
		return Token{}, err
	}

	return token, nil
}
