package action

import (
	"github.com/gin-gonic/gin"

	"jammming/auth"

	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	baseURL = "https://api.spotify.com/v1/"
)

// Search Spotify

type searchResponse struct {
	Tracks struct {
		Items []struct {
			ID      string `json:"id"`
			URI     string `json:"uri"`
			Name    string `json:"name"`
			Artists []struct {
				Name string `json"name"`
			}
			Album struct {
				Images []struct {
					URL string `json:"url"`
				}
			}
		}
	}
}

type track struct {
	ID     string `json"id"`
	URI    string `json"uri"`
	Name   string `json"name"`
	Artist string `json"artist"`
	Image  string `json"image"`
}

type tracks struct {
	Tracks []track
}

func SearchSpotify(c *gin.Context) {
	if auth.Access.Token == "" {
		c.JSON(http.StatusOK, gin.H{"error": "No access token"})
		return
	}

	endpoint := fmt.Sprintf("%ssearch?q=%s&type=track&limit=15", baseURL, url.QueryEscape(c.Query("searchTerm")))

	client := &http.Client{}

	req, _ := http.NewRequest(http.MethodGet, endpoint, nil)
	req.Header.Add("Authorization", "Bearer "+auth.Access.Token)

	resp, _ := client.Do(req)

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusOK, gin.H{"error": fmt.Sprintf("Spotify API returned status code: %d", resp.StatusCode)})
		return
	}

	var respJSON searchResponse

	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &respJSON)

	var results tracks

	for _, respTrack := range respJSON.Tracks.Items {
		var t track

		t.ID = respTrack.ID
		t.URI = respTrack.URI
		t.Name = respTrack.Name
		if len(respTrack.Artists) > 1 {
			var artists []string
			for _, a := range respTrack.Artists {
				artists = append(artists, a.Name)
			}
			t.Artist = strings.Join(artists, ", ")
		} else {
			t.Artist = respTrack.Artists[0].Name
		}

		t.Image = respTrack.Album.Images[0].URL

		results.Tracks = append(results.Tracks, t)
	}
	c.JSON(http.StatusOK, results)
}

// Create Playlist

type playlistInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var playlist playlistInfo

func CreatePlaylist(c *gin.Context) {
	if auth.Access.Token == "" {
		c.JSON(http.StatusOK, gin.H{"error": "No access token"})
		return
	}

	if auth.User.ID == "" {
		c.JSON(http.StatusOK, gin.H{"error": "No userID"})
	}

	endpoint := fmt.Sprintf("%susers/%s/playlists", baseURL, auth.User.ID)

	jsonData, _ := c.GetRawData()

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(jsonData))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+auth.Access.Token)

	resp, _ := client.Do(req)

	if resp.StatusCode != http.StatusCreated {
		c.JSON(http.StatusOK, gin.H{"error": fmt.Sprintf("Spotify API returned status code: %d", resp.StatusCode)})
		return
	}

	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &playlist)

	c.JSON(http.StatusOK, gin.H{"success": playlist.ID})

}

// Add Tracks to Playlist

func AddTracks(c *gin.Context) {
	if auth.Access.Token == "" {
		c.JSON(http.StatusOK, gin.H{"error": "No access token"})
		return
	}

	if playlist.ID == "" {
		c.JSON(http.StatusOK, gin.H{"error": "No Playlist ID"})
		return
	}

	endpoint := fmt.Sprintf("%splaylists/%s/tracks", baseURL, playlist.ID)

	jsonData, _ := c.GetRawData()

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(jsonData))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+auth.Access.Token)

	resp, _ := client.Do(req)

	if resp.StatusCode != http.StatusCreated {
		c.JSON(http.StatusOK, gin.H{"error": fmt.Sprintf("Spotify API returned status code: %d", resp.StatusCode)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Tracks successfully added to playlist: " + playlist.Name})
}
