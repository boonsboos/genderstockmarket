package routes

import (
	"math/rand"
	"net/http"
	util "spectrum300/Util"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixMilli()))
}

const allowedChars = "_-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Signup(context *gin.Context) {

	// generate random state string of at least 32 characters
	state = util.GenerateRandomString(32)

	// redirect to github OAuth portal
	context.Redirect(http.StatusPermanentRedirect,
		"https://github.com/login/oauth/authorize?client_id="+
			util.Options.GithubID+
			"&redirect_uri="+
			"http://localhost:8100/submit"+ // FIXME: prod URL
			"&scope=read:user"+
			"&state="+string(state)+
			"&allow_signup=true",
	)

}

func SubmitSignup(context *gin.Context) {
	authCode, found := context.GetPostForm("code")
	if !found || authCode == "" {
		context.String(http.StatusForbidden, "failed to log in, try again")
		time.Sleep(3 * time.Second)
		context.Redirect(http.StatusPermanentRedirect, "/login")
	}

	// use code to make request for username
	// create new player
	// save player to database

	// create a new oauth client

	// redirect to success (play?) page
}
