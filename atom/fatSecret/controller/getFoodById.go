package atom_fatSecret_controller

import (
	"log"
	"net/http"
	atom_fatSecret "optimach_service/atom/fatSecret"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFoodById(ctx *gin.Context) {
	foodId := ctx.Query("food_id")
	foodIdInt, _ := strconv.Atoi(foodId)
	format := ctx.DefaultQuery("format", "json")

	if foodId == "" {
		log.Println("[atom][fatSecret][controller][GetFoodById] missing food_id parameter")

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "food_id parameter is required",
		})
		return
	}

	results, _, err := atom_fatSecret.GetFoodByIdUseCase(foodIdInt, format)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "failed to get food by id",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Success get fatsecret food by id!",
		"results": results,
	})
}
