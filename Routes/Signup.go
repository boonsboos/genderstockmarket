package routes

import (
	ctx "context"
	"log"
	"math/rand"
	"net/http"
	auth "spectrum300/Auth"
	entities "spectrum300/Entities"
	util "spectrum300/Util"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixMilli()))
}

func Signup(context *gin.Context) {

	// generate random state string of at least 32 characters
	state := auth.GenerateRandomString(32)

	// FIXME: prod domain
	context.SetCookie("oauthState", state, 60, "", "localhost", true, true)
	// NB: should probably be saved server side

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

func loginFail(context *gin.Context) {
	context.String(http.StatusForbidden, "failed to log in, try again")
	time.Sleep(3 * time.Second)
	context.Redirect(http.StatusPermanentRedirect, "/login")
}

type SignupResponse struct {
	ID     string `json:"id"`
	Secret string `json:"secret"`
}

func SubmitSignup(context *gin.Context) {
	// get auth code to make requests to github
	authCode, found := context.GetQuery("code")
	if !found || authCode == "" {
		loginFail(context)
		log.Println("code param not found")
		return
	}

	// // get state string
	// state, found := context.GetQuery("state")
	// if !found || state == "" {
	// 	loginFail(context)
	// 	log.Println("state param not found")
	// 	return
	// }

	// stateCookie, err := context.Cookie("oauthState")
	// if err != nil {
	// 	loginFail(context)
	// 	log.Println("cookie oauthState not found")
	// 	return
	// }

	// // verify if state sent by github is the same as
	// // the state we saved as a cookie
	// if state != stateCookie {
	// 	loginFail(context)
	// 	log.Println("State does not match")
	// 	return
	// }

	token, err := util.GetUserAccessToken(authCode)
	if token == "" || err != nil {
		context.String(http.StatusBadRequest, "something went wrong. try again later!")
		return
	}

	username, err := util.GetGithubUsername(token)
	if username == "" || err != nil {
		context.String(http.StatusBadRequest, "something went wrong. try again later!")
		return
	}

	user, err := auth.ClientStore.GetByID(ctx.Background(), username)
	if err != nil {
		log.Println(err.Error(), "Attempting to create new client...")
	}

	if user.Domain == "" {
		err = entities.SaveNewPlayer(username)
		if err != nil {
			log.Println(err.Error())
			return
		}

		user, err = auth.ClientStore.GetByID(ctx.Background(), username)
		if err != nil {
			log.Println("Player client still not found by ID:", err.Error())
			return
		}
	}

	context.JSON(http.StatusOK, SignupResponse{
		ID:     username,
		Secret: user.Secret,
	})
}
