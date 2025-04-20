package main

import (
	"github.com/gin-gonic/gin"
	
	"jammming/auth"
	
)

func main() {
	router := gin.Default()

	router.GET("auth/login", auth.RedirectToAuthURL)
	router.GET("/auth/callback", auth.HandleAuthCallback)

	router.Run("localhost:5000")
}

