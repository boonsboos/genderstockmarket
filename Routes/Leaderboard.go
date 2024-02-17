package routes

import "github.com/gin-gonic/gin"

// provides endpoints for leaderboards

// GET /leaderboards
//
// returns the endpoints for the leaderboards we have.
func AllLeaderboards(context *gin.Context) {

}

// GET /leaderboard/networth
//
// returns the top 50 highest net worth players
func NetWorthLeaderboard(context *gin.Context) {

}

// GET /leaderboard/trades
//
// returns the top 50 busiest traders
func TradesLeaderboard(context *gin.Context) {

}
