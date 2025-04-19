package main

import (
	"github.com/gin-gonic/gin"
	
	"jammming/auth"
	
	"net/http"
	"encoding/json"
	"os"
	"log"
	"strconv"
)

func main() {
	router := gin.Default()

	router.POST("/login")
	router.GET("/auth")
}

