package main

import (
	auth "spectrum300/Auth"
	database "spectrum300/Database"
	routes "spectrum300/Routes"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	auth.InitAuthServer(r)
	routes.RegisterAll(r)

	r.Run(":8100")

	defer database.DatabaseConnection.Close()
}
