package playback

import (
	"github.com/gin-gonic/gin"
	"jammming/auth"

	"net/http"
	"fmt"
	"bytes"
)

const baseURL = "https://api.spotify.com/v1/me/player"

func TransferPlayback(c *gin.Context) {
	if auth.Access.Token == "" {
		c.JSON(http.StatusOK, gin.H{"error": "No access token"})
		return
	}

	jsonData, _ := c.GetRawData()

	client := &http.Client{}
	
	req, _ := http.NewRequest(http.MethodPut, baseURL, bytes.NewReader(jsonData))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer " + auth.Access.Token)

	resp, _ := client.Do(req)

	if resp.StatusCode != http.StatusNoContent {
		c.JSON(http.StatusOK, gin.H{"error": fmt.Sprintf("Spotify returned the status code: %d", resp.StatusCode)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Playback Transferred to SDK"})
}

func PlayTrack(c *gin.Context) {
	if auth.Access.Token == "" {
		c.JSON(http.StatusOK, gin.H{"error": "No access token"})
		return
	}

	endpoint := fmt.Sprintf("%s/play", baseURL)

	jsonData, _ := c.GetRawData()

	client := &http.Client{}

	req, _ := http.NewRequest(http.MethodPut, endpoint, bytes.NewReader(jsonData))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer " + auth.Access.Token)

	resp, _ := client.Do(req)

	if resp.StatusCode != http.StatusNoContent {
		c.JSON(http.StatusOK, gin.H{"error": fmt.Sprintf("Spotify returned the status code: %d", resp.StatusCode)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Track now playing"})
}

func PausePlayback(c *gin.Context) {
	if auth.Access.Token == "" {
		c.JSON(http.StatusOK, gin.H{"error": "No access token"})
	}

	endpoint := fmt.Sprintf("%s/pause", baseURL)

	client := &http.Client{}

	req, _ := http.NewRequest(http.MethodPut, endpoint, nil)
	req.Header.Add("Authorization", "Bearer " + auth.Access.Token)

	resp, _ := client.Do(req)

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusOK, gin.H{"error": fmt.Sprintf("Spotify returned the status code: %d", resp.StatusCode)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Playback paused"})
}
