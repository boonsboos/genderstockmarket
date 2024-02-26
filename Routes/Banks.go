package routes

import (
	"log"
	"net/http"
	auth "spectrum300/Auth"
	entities "spectrum300/Entities"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

// banks are where you get interest on your money while you're not trading!
// while it is slow and takes a long time to get more money (an ingame year!)
// it will get you a nice little pot of money where you save

// GET /banks
//
// returns a list of all banks
func Banks(context *gin.Context) {

}

// GET /bank/:id
//
// gets the info of a bank by its ID
func BankInfo(context *gin.Context) {

}

// GET /bank/:id/open
//
// open an account at bank, paying the fee
func OpenBankAccount(context *gin.Context) {

}

// GET /bank/accounts/:id/close
//
// close your account
func CloseBankAccount(context *gin.Context) {

}

// GET /bank/accounts
//
// check the balance of your bank accounts, also shows the account numbers
func BankAccountsList(context *gin.Context) {
	token, err := auth.ValidateBearerToken(context.Request.Header.Get("Authorization"))
	if err != nil {
		context.JSON(http.StatusUnauthorized, auth.Unauthorized)
		return
	}

	player, err := auth.ClientStore.GetByNumericID(token.ID)
	if err != nil {
		log.Println("Failed to ")
		context.JSON(http.StatusNotFound, auth.AuthError{
			Kind:    "not found",
			Message: "try again later",
		})
		return
	}

	accounts, err := entities.GetPlayerBankAccounts(token.ID)
	if err != nil {
		log.Println("Failed to get player's bank accounts")
		context.JSON(http.StatusNotFound, auth.AuthError{
			Kind:    "not found",
			Message: "try again later",
		})
		return
	}

	context.JSON(http.StatusOK, struct {
		Username string                 `json:"username"`
		Accounts []entities.BankAccount `json:"accounts"`
	}{
		player.ID,
		accounts,
	})
}

// GET /bank/account/:id/withdraw?amount
//
// withdraws amount to your "wallet" from the specified account
// if amount is not specified, return an error.
func WithdrawFromBank(context *gin.Context) {
	token, err := auth.ValidateBearerToken(context.Request.Header.Get("Authorization"))
	if err != nil {
		context.JSON(http.StatusUnauthorized, auth.Unauthorized)
		return
	}

	account, _ := strconv.Atoi(context.Param("id"))

	belongs, err := entities.CheckAccountBelongsToPlayer(token.ID, account)
	if err != nil {
		log.Println("Failed to withdraw while checking if account belongs to player: ", token.ID, err.Error())
	}

	if !belongs {
		context.JSON(http.StatusUnauthorized, auth.Unauthorized)
		return
	}

	amount := context.Query("amount")
	if amount == "" {
		context.JSON(http.StatusBadRequest, ParameterMissing)
		return
	}

	amnt, err := decimal.NewFromString(amount)
	if err != nil {
		log.Println("Failed to create decimal from amount while withdrawing for player:", token.ID, err.Error())
		context.JSON(http.StatusInternalServerError, auth.InternalServerError)
		return
	}

	entities.Withdraw(token.ID, account, amnt)

	stats, err := entities.GetPlayerProfile(token.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, auth.InternalServerError)
		return
	}

	context.JSON(http.StatusOK, stats)
}

// GET /bank/account/:id/deposit?amount
//
// deposits amount from your "wallet" to the specified account
// if amount is not specified, return an error.
func DepositInBank(context *gin.Context) {

}

// GET /bank/account/:id/newsletter
//
// returns the bank newsletter (only available if account holder)
func BankNewsletter(context *gin.Context) {

}

// GET /bank/account/transfer/:id?amount
//
// transfers amount to account id
func BankTransfer(context *gin.Context) {

}
