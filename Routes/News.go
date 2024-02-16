package routes

import "github.com/gin-gonic/gin"

// news sources about companies

// GET /papers
//
// returns a list of all newspapers you can subscribe to
func Newspapers(context *gin.Context) {

}

// GET /papers/:id
//
// returns the news in that paper if subscribed
func Paper(context *gin.Context) {

}

// GET /papers/:id/subscribe
//
// subscribes to the paper
func SubToPaper(context *gin.Context) {

}

// GET /papers/:id/unsubscribe
//
// unsubscribes from the paper
func UnsubFromPaper(context *gin.Context) {

}
