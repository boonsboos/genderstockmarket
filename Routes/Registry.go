package routes

import (
	"log"

	"github.com/gin-gonic/gin"
)

func RegisterAll(router *gin.Engine) {
	router.GET("/login", Signup) // TODO remove because this is for frontend
	router.GET("/submit", SubmitSignup)

	// banks
	router.GET("/bank/accounts", BankAccountsList)
	router.GET("/bank/account/:id/withdraw", WithdrawFromBank)

	// player
	router.GET("/me/total", PlayerTotalBalance)

	// auth
	router.GET("/token", TokenRequest)

	log.Println("API routes OK")
}
