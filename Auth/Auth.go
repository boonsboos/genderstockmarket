package auth

import (
	"log"
	"net/http"
	"sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	_ "github.com/jackc/pgx/v5/stdlib"
	oauthpg "github.com/vgarvardt/go-oauth2-pg/v4"
	"github.com/vgarvardt/go-pg-adapter/pgx4adapter"
)

var authServer server.Server

// https://github.com/go-oauth2/oauth2 readme
func InitAuthServer(router *gin.Engine) {

	pgxConn, err := sql.Open("")
	if err != nil {
		log.Println("Failed to connect to database:", err.Error())
	}

	manager := manage.NewDefaultManager()

	// use PostgreSQL token store with pgx.Connection adapter
	adapter := pgx4adapter.NewConn(pgxConn)
	tokenStore, _ := oauthpg.NewTokenStore(adapter, oauthpg.WithTokenStoreGCInterval(time.Minute))

	// client memory store
	clientStore := store.NewClientStore()
	clientStore.Set("0", &models.Client{
		ID:     "0",
		Secret: "1",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientStore)

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

	// authorize
	router.GET("/authorize", func(context *gin.Context) {
		err := authServer.HandleAuthorizeRequest(context.Writer, context.Request)
		if err != nil {
			context.String(http.StatusBadRequest, err.Error())
		}
	})

	// getting token
	router.GET("/token", func(context *gin.Context) {
		authServer.HandleTokenRequest(context.Writer, context.Request)
	})
}
