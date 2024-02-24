package routes

import (
	"net/http"
	auth "spectrum300/Auth"
	"time"

	"github.com/gin-gonic/gin"
)

// GET /token
//
// semi-oauth token request
func TokenRequest(context *gin.Context) {

	client, secret, err := auth.ValidateTokenRequest(context.Request.URL.Query())
	if err != nil {
		context.JSON(http.StatusBadRequest, auth.AuthError{
			Kind:    "bad request",
			Message: err.Error(),
		})
		return
	}

	token, err := auth.CreateNewToken(client, secret)
	if err != nil {
		context.JSON(http.StatusInternalServerError, auth.InternalServerError)
		return
	}

	// standard oauth2 token response
	response := struct {
		Scope     string        `json:"scope"`
		ExpiresIn time.Duration `json:"expires_in"`
		Code      string        `json:"access_code"`
		TokenType string        `json:"token_type"`
	}{
		"read",
		token.ExpiresAt.Sub(token.CreatedAt) / time.Second,
		token.Code,
		"Bearer",
	}

	context.JSON(http.StatusOK, response)
}
