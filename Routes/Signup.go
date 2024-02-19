package routes

import (
	"log"
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

	// generate random state string of at least 16 characters
	state := make([]byte, 32)

	for character := range state {
		state[character] = allowedChars[rand.Int63()%int64(len(allowedChars))]
	}

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

	log.Println(context.Request)
	// oauth callback
	// create new player
	// create new auth client
	// redirect user to signup complete page
}
