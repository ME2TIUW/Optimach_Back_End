package atom_auth

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	utils_fatsecret "optimach_service/utils/fatsecret"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func executeFatSecretRequest(apiMethod string, params url.Values) (map[string]interface{}, int, error) {
	token, err := utils_fatsecret.GetValidToken()
	if err != nil {
		log.Println("[atom][fatsecret] Failed to get access token:", err)
		return nil, http.StatusInternalServerError, err
	}

	params.Set("method", apiMethod)
	params.Set("format", "json")

	req, err := http.NewRequest("GET", os.Getenv("FS_BASE_URL"), nil)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.URL.RawQuery = params.Encode()

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, http.StatusBadGateway, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return result, resp.StatusCode, nil
}

func SearchFood(ctx *gin.Context) {
	query := ctx.Query("query")
	pageNumber := ctx.DefaultQuery("page_number", "0")
	maxResults := ctx.DefaultQuery("max_results", "16")

	if query == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Search query is required",
		})
		return
	}

	params := url.Values{}
	params.Set("search_expression", query)
	params.Set("page_number", pageNumber)
	params.Set("max_results", maxResults)

	data, statusCode, err := executeFatSecretRequest("foods.search", params)
	if err != nil {
		ctx.JSON(statusCode, gin.H{"status": statusCode, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Success",
		"results": data,
	})
}

func GetFoodDetailById(ctx *gin.Context) {
	foodId := ctx.Query("food_id")

	if foodId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "food_id is required",
		})
		return
	}

	params := url.Values{}
	params.Set("food_id", foodId)

	data, statusCode, err := executeFatSecretRequest("food.get.v2", params)
	if err != nil {
		ctx.JSON(statusCode, gin.H{"status": statusCode, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Success",
		"results": data,
	})
}
