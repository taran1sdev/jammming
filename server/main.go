package main

import (
	"github.com/gin-gonic/gin"

	"jammming/action"
	"jammming/auth"
	"jammming/playback"

	"net/http"
)

func main() {
	router := gin.Default()

	// Authorization Endpoints
	router.GET("auth/login", auth.RedirectToAuthURL)
	router.GET("/auth/callback", auth.HandleAuthCallback)
	router.GET("/auth", auth.GetAuthenticated)

	// User Info Endpoints
	router.GET("/access", getAccessToken)
	router.GET("/userId", getUserID)

	// Spotify actions
	router.GET("/search", action.SearchSpotify)

	router.POST("/createPlaylist", action.CreatePlaylist)
	router.POST("/addTracks", action.AddTracks)

	router.POST("/transferPlayback", playback.TransferPlayback)
	router.POST("/playTrack", playback.PlayTrack)

	router.GET("/pause", playback.PausePlayback)

	router.Run("localhost:5000")
}

func getAccessToken(c *gin.Context) {
	if auth.Access.Token == "" {
		c.Redirect(http.StatusFound, "http://localhost:5000/auth/login")
	}

	c.JSON(http.StatusOK, gin.H{"access_token": auth.Access.Token})
}

func getUserID(c *gin.Context) {
	if auth.User.ID == "" {
		c.JSON(http.StatusOK, gin.H{"error": "No userID"})
	}

	c.JSON(http.StatusOK, gin.H{"user_id": auth.User.ID})
}
