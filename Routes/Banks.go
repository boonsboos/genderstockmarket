package routes

import "github.com/gin-gonic/gin"

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
func BankBalance(context *gin.Context) {

}

// GET /bank/account/:id/withdraw?amount
//
// withdraws amount to your "wallet" from the specified account
// if amount is not specified, return an error.
func WithdrawFromBank(context *gin.Context) {

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
