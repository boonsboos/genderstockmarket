package routes

import "github.com/gin-gonic/gin"

// routes to get info from a company

// GET /companies
//
// returns a list of all companies listed on the stock market
// both name and id.
func Companies(context *gin.Context) {

}

// GET /companies/:id
//
// id - either the numeric ID or the stock name
//
// returns information about a company, eg industry, stock listing
// other options for equities they have like obligations,
func Company(context *gin.Context) {

}

// GET /companies/:id/balance
//
// returns the balance sheet of the company
func CompanyBalance(context *gin.Context) {

}

// GET /companies/:id/newsletter
//
// returns the company newsletter (only available if a shareholder)
func CompanyNewsletter(context *gin.Context) {

}
