package utils_fatsecret

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

// Struct to hold the token and its expiration time
type FatSecretToken struct {
	AccessToken string
	Expiry      time.Time
	mu          sync.RWMutex // Mutex to prevent race conditions
}

var globalToken = &FatSecretToken{}

func GetValidToken() (string, error) {
	globalToken.mu.RLock()
	if globalToken.AccessToken != "" && time.Now().Add(60*time.Second).Before(globalToken.Expiry) {
		defer globalToken.mu.RUnlock()
		return globalToken.AccessToken, nil
	}
	globalToken.mu.RUnlock()

	return fetchNewToken()
}

func fetchNewToken() (string, error) {
	globalToken.mu.Lock()
	defer globalToken.mu.Unlock()

	if globalToken.AccessToken != "" && time.Now().Add(60*time.Second).Before(globalToken.Expiry) {
		return globalToken.AccessToken, nil
	}

	log.Println("[utils][fatsecret] Token expired or missing. Fetching new token...")

	clientID := os.Getenv("FS_CLIENT_ID")
	clientSecret := os.Getenv("FS_CLIENT_SECRET")
	tokenURL := os.Getenv("FS_ACCESS_TOKEN_URL")

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("scope", "basic")

	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var tokenResponse struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return "", err
	}

	globalToken.AccessToken = tokenResponse.AccessToken
	globalToken.Expiry = time.Now().Add(time.Duration(tokenResponse.ExpiresIn) * time.Second)

	log.Printf("[utils][fatsecret] New token acquired. Expires in: %d seconds", tokenResponse.ExpiresIn)

	return globalToken.AccessToken, nil
}
