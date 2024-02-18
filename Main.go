package main

import (
	auth "spectrum300/Auth"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	auth.InitAuthServer(r)

	r.Run(":8100")
}
