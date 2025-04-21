package auth

import(
	"fmt"
	"strings"
	"net/http"
	"net/url"
	"encoding/base64"
	"encoding/json"
	"log"
	"io"
	"os"
	"errors"

	"github.com/gin-gonic/gin"
)

const (
	frontend = "http://localhost:3000/"

	authEndpoint = "https://accounts.spotify.com/authorize"
	accessTokenEndpoint = "https://accounts.spotify.com/api/token"
	userEndpoint = "https://api.spotify.com/v1/me"

	redirectUri = "http://localhost:5000/auth/callback"
)

// datatype for the response from the access /api/token endpoint
type access struct {
	Token        string `json:"access_token"`
	Type  	     string `json:"token_type"`
	Expires      int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope	     string `json:"scope"`
}

type user struct {
	Name	string `json:"display_name"`
	ID	string `json:"id"` 
}

var (
	Access access
	User user

	clientId string = os.Getenv("CLIENT_ID") 
	clientSecret string = os.Getenv("CLIENT_SECRET")
)

// Handles the redirect to spotify account login
func RedirectToAuthURL(c *gin.Context) {
	scopes := []string{
	"streaming",
	"user-read-private",
	"user-read-email",
	"playlist-modify-public",
	"user-read-private",
	"playlist-modify-private",
}


	authURL := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&scope=%s&response_type=code&show_dialog=true", authEndpoint, clientId, redirectUri, strings.Join(scopes[:], "%20"))

	c.Redirect(http.StatusFound, authURL)
}

// Function to return url form data for the /api/token request
func getUrlFormData(code string) url.Values {
	data := url.Values{}
	data.Add("grant_type", "authorization_code")
	data.Add("code", code)
	data.Add("redirect_uri", redirectUri)
	data.Add("client_id", clientId)
	data.Add("client_secret", clientSecret)

	return data
}

func getUserInfo() error {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, userEndpoint, nil)
	req.Header.Add("Authorization", "Bearer " + Access.Token)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		json.Unmarshal(body, &User)
		return nil
	} else {
		return errors.New("Bad response when requesting user data: " + http.StatusText(resp.StatusCode))

	}
}

// Handles the callback containing auth code and makes request to /api/token and stores response in the Access variable
func HandleAuthCallback(c *gin.Context) {
	authCode := c.Query("code")
	if authCode == "" {
		c.String(http.StatusBadRequest, "Auth code not found in the response")
		return
	}


	data := getUrlFormData(authCode)
	
	authHeaderString := clientId + ":" + clientSecret

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, accessTokenEndpoint, strings.NewReader(data.Encode()))
	req.Header.Add("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte(authHeaderString)))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		json.Unmarshal(body, &Access)
		if err = getUserInfo(); err != nil {
			log.Fatal("Error getting user info")
		}
		c.Redirect(http.StatusFound, frontend)
	} else {
		log.Fatalf("Request returned Status %d: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

}
