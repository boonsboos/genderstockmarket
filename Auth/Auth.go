package auth

import (
	"log"
	database "spectrum300/Database"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	oauthpg "github.com/vgarvardt/go-oauth2-pg/v4"
	"github.com/vgarvardt/go-pg-adapter/pgx4adapter"
)

var AuthServer server.Server
var ClientStore SpectrumClientStore

// https://github.com/go-oauth2/oauth2 readme
func InitAuthServer(router *gin.Engine) {
	// database stores
	adapterConn := pgx4adapter.NewPool(&database.DatabaseConnection)

	tokenStore, err := oauthpg.NewTokenStore(adapterConn, oauthpg.WithTokenStoreGCInterval(time.Minute))
	if err != nil {
		log.Fatal("Failed to created token store:", err.Error())
	}
	log.Println("Token store OK")

	ClientStore, err := NewClientStore()
	if err != nil {
		log.Fatal("Failed to created client store:", err.Error())
	}
	log.Println("Client store OK")

	manager := manage.NewDefaultManager()
	manager.MapClientStorage(ClientStore)
	manager.MapTokenStorage(tokenStore)

	authServer := server.NewDefaultServer(manager)
	authServer.SetAllowGetAccessRequest(true)
	authServer.SetClientInfoHandler(server.ClientFormHandler)

	authServer.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	authServer.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	// getting token
	router.GET("/token", func(context *gin.Context) {
		authServer.HandleTokenRequest(context.Writer, context.Request)
	})

	log.Println("Auth server OK")
}
