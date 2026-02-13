package atom_fatSecret

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var cache = &tokenCacheModel{
	httpClient: &http.Client{Timeout: 10 * time.Second},
}

func (tc *tokenCacheModel) fetchNewToken() error {
	log.Println("Fetch new fatsecret token")

	clientID := os.Getenv("FS_CLIENT_ID")
	clientSecret := os.Getenv("FS_CLIENT_SECRET")
	tokenUrl := os.Getenv("FS_ACCESS_TOKEN_URL")

	if clientID == "" || clientSecret == "" {
		log.Println("FATSECRET_CLIENT_ID or FATSECRET_CLIENT_SECRET is not set in environment")
		return errors.New("server configuration error")
	}

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("scope", "basic")

	req, err := http.NewRequest("POST", tokenUrl, strings.NewReader(data.Encode()))
	if err != nil {
		log.Println("[usecase][get_token] error creating request:", err.Error())
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := tc.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("FatSecret API returned status %d: %s", resp.StatusCode, string(bodyBytes))

		log.Println(errMsg)
		return errors.New(errMsg)
	}

	var tokenResponse fatSecretTokenResponseModel
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		log.Println("[usecase][get_token] error decoding response:", err.Error())
		return err
	}

	tc.Token = tokenResponse.Access_Token
	tc.ExpiryTime = time.Now().Add(time.Duration(tokenResponse.Expires_In-300) * time.Second)
	log.Println("Successfully fetched and cached new FatSecret token")
	return nil
}

func GetValidToken() (string, error) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	log.Println("cache expiry time", cache.ExpiryTime)
	log.Println("ctime.now", time.Now())

	if time.Now().Before(cache.ExpiryTime) {
		log.Println("Returning cached FatSecret token")
		return cache.Token, nil
	}

	if err := cache.fetchNewToken(); err != nil {
		return "", err
	}

	return cache.Token, nil
}

var searchHttpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func SearchFoodUseCase(query, pageNumber, maxResults string) (interface{}, bool, error) {

	token, err := GetValidToken()

	if err != nil {
		return nil, false, errors.New("couldn't authenticate with food service")
	}

	fatSecretURL := "https://platform.fatsecret.com/rest/foods/search/v1"

	req, err := http.NewRequest("GET", fatSecretURL, nil)
	if err != nil {
		return nil, false, errors.New("failed to create search request")
	}

	data := url.Values{}
	data.Set("search_expression", query)
	data.Set("page_number", pageNumber)
	data.Set("max_results", maxResults)
	data.Set("format", "json")
	req.URL.RawQuery = data.Encode()

	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := searchHttpClient.Do(req)
	if err != nil {
		return nil, false, errors.New("failed to contact food service")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("FatSecret API error. Status: %d, Body: %s", resp.StatusCode, string(bodyBytes))
		log.Println(errMsg)

		return nil, false, errors.New("food service returned an error")
	}

	var results interface{}
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, false, errors.New("failed to parse food service response")
	}

	return results, true, nil
}

func GetFoodByIdUseCase(foodId int, format string) (interface{}, bool, error) {
	token, err := GetValidToken()

	if err != nil {
		return nil, false, errors.New("couldn't authenticate with food service")
	}

	fatSecretURL := "https://platform.fatsecret.com/rest/food/v1"

	req, err := http.NewRequest("GET", fatSecretURL, nil)

	if err != nil {
		return nil, false, errors.New("failed to create get food by id request")
	}

	data := url.Values{}
	data.Set("food_id", strconv.Itoa(foodId))
	data.Set("format", format)
	req.URL.RawQuery = data.Encode()

	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := searchHttpClient.Do(req)
	if err != nil {
		return nil, false, errors.New("failed to contact food service")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("FatSecret API error. Status: %d, Body: %s", resp.StatusCode, string(bodyBytes))
		log.Println(errMsg)

		return nil, false, errors.New("food service returned an error")
	}

	var results interface{}
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, false, errors.New("failed to parse food service response")
	}

	return results, true, nil
}
