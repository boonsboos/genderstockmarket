package routes

import "github.com/gin-gonic/gin"

// this package has things related to firms
// firms are like clans, you can join and work for them
// you can also create your own firm

// GET /firms
//
// returns a list of all firms
func Firms(context *gin.Context) {

}

// GET /firm/:id
//
// gets information about a firm by its id
// also should show all members' name
func Firm(context *gin.Context) {

}

// GET /firm/:id/join
//
// join a firm if allowed to join
func FirmJoin(context *gin.Context) {

}

// GET /firm/:id/leave
//
// leave a firm if you are in it
func FirmLeave(context *gin.Context) {

}
