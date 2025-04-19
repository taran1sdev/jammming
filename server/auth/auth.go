package auth

import(
	"fmt"
	"strings"
	
	"github.com/gin-gonic/gin"
)

const (
	clientId = "985aa6f5dba84edc8e92f5fda306d18e"
	clientSecret = "3d5ae77c544e48ba97a82e9201c50fc4"

	authEndpoint = "https://accounts.spotify.com/authorize"
	accessTokenEndpoint = "https://accounts.spotify.com/api/token"

	redirectUri = "http://localhost:3000/auth/callback"
)

scopes := []string{}{
	"streaming",
	"user-read-private",
	"user-read-email"
}

const loginURL := fmt.Sprintf("%v?client_id=%v&redirect_uri=%v&scope=%v&response_type=code&show_dialog=true"authEndpoint, clientId, redirectUri, strings.Join(scopes[:], "%20")) 

func RedirectToAuthURL(c *gin.Context) {
	authURL := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&scope=%s&response_type=code&show_dialog=true"authEndpoint, clientId, redirectUri, strings.Join(scopes[:], "%20"))

	c.Redirect(http.StatusFound, authURL)
}

func HandleAuthCallback(c *gin.Context) {
	authCode := c.Query("code")
	if authCode == "" {
		c.String(http.StatusBadRequest, "Auth code not found in the response")
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"auth": authCode}
}
