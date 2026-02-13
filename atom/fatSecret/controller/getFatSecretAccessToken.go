package atom_fatSecret_controller

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	atom_auth "optimach_service/atom/auth"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetFatSecretAccessToken(ctx *gin.Context) {
	clientId := os.Getenv("FS_CLIENT_ID")
	clientSecret := os.Getenv("FS_CLIENT_SECRET")
	tokenUrl := os.Getenv("FS_ACCESS_TOKEN_URL")

	if clientId == "" || clientSecret == "" {
		log.Println("[atom][auth][controller][GetFatSecretAccessToken] missing client ID or client secret")

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status":  500,
			"Message": "Server configuration error",
		})
	}

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	data.Set("scope", "basic")

	req, err := http.NewRequest("POST", tokenUrl, strings.NewReader(data.Encode()))

	if err != nil {
		log.Println("[atom][auth][controller][GetFatSecretAccessToken] error creating request:", err.Error())

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Error creating request",
		})
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500",
			"message": "Failed to Fetch Access Token",
		})
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  resp.StatusCode,
			"message": "FatSecret API returned an error",
		})
		return
	}

	// decode JSON response and send only the token back

	var tokenresponse atom_auth.FatSecretTokenResponseModel

	if err := json.NewDecoder(resp.Body).Decode(&tokenresponse); err != nil {
		log.Println("[atom][auth][controller][GetFatSecretAccessToken] error decoding response:", err.Error())

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500",
			"message": "Error decoding response",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":       200,
		"access_token": tokenresponse.Access_Token,
		"expires_in":   tokenresponse.Expires_In,
	})
}
