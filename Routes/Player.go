package routes

import (
	"net/http"
	auth "spectrum300/Auth"
	entities "spectrum300/Entities"

	"github.com/gin-gonic/gin"
)

// routes pertaining to the player

// GET /me
//
// shows an overview of the player profile with (the resources for) statistics.
// the resources should be calculated into stats client side
func PlayerStats(context *gin.Context) {
	// token, err := auth.ValidateBearerToken(context.Request.Header.Get("Authorization"))
	// if err != nil {
	// 	context.String(http.StatusUnauthorized, "no permission")
	// 	return
	// }

	// get player ID by their username

	// context.JSON(http.StatusOK, entities.PlayerInfo{
	// 	token.GetClientID(),
	// 	,
	// })
}

// GET /me/total
//
// shows the player their total balance
func PlayerTotalBalance(context *gin.Context) {
	token, err := auth.ValidateBearerToken(context.Request.Header.Get("Authorization"))
	if err != nil {
		context.JSON(http.StatusUnauthorized, auth.Unauthorized)
		return
	}

	number, err := entities.GetPlayerTotalBalance(token.ID)

	context.JSON(http.StatusOK, struct {
		Balance string `json:"balance"`
	}{
		number.String(),
	})
}
