package routes

import "github.com/gin-gonic/gin"

// contains routes for stuff to do with the market

// GET /market
//
// returns the endpoints for the different types of equity
func Market(context *gin.Context) {

}

// GET /market/stocks
//
// returns all currently available stock and their price
func StockMarket(context *gin.Context) {

}

// GET /market/stocks/:id?time=(\d(d|w|m|y))
//
// returns (historical) pricing data for a certain stock
// default should be one day
func HistoricalStockPrices(context *gin.Context) {

}

// GET /market/stocks/:id/purchase?amount
//
// amount - 1 if omitted.
// purchases the amount of the stock
func PurchaseStock(context *gin.Context) {

}

// GET /market/stocks/:id/sell?amount
//
// amount - 1 if omitted.
// purchases the amount of the provided stock
func SellStock(context *gin.Context) {

}

// GET /market/stocks/:id/call?time=(\d(d|w|m|y))
//
// buys a call option in the provided stock during the selected time frame
// (expects the stock to go down)
func BuyStockCall(context *gin.Context) {

}

// GET /market/stocks/:id/put?time=(\d(d|w|m|y))
//
// buys a put option in the provided stock during the selected time frame
// (expects the stock to go up)
func BuyStockPut(context *gin.Context) {

}
