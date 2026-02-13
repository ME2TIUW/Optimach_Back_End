package atom_fatSecret_controller

import (
	"net/http"
	atom_fatSecret "optimach_service/atom/fatSecret"

	"github.com/gin-gonic/gin"
)

func GetFoodSearch(ctx *gin.Context) {
	query := ctx.Query("query")
	pageNumber := ctx.DefaultQuery("page_number", "0")
	maxResults := ctx.DefaultQuery("max_results", "20")

	if query == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "query parameter is required",
		})
		return
	}

	results, _, err := atom_fatSecret.SearchFoodUseCase(query, pageNumber, maxResults)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "failed to search food",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "food search successful",
		"results": results,
	})
}
