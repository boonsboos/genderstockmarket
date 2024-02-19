package routes

import (
	"log"

	"github.com/gin-gonic/gin"
)

func RegisterAll(router *gin.Engine) {
	router.GET("/login", Signup)
	router.POST("/submit", SubmitSignup)

	log.Println("API routes OK")
}
