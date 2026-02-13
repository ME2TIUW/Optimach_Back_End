package atom_masterdata_food

import (
	"net/http"
	atom_masterdata_food "optimach_service/atom/masterdata/food"

	"github.com/gin-gonic/gin"
)

func PutDeleteFood(ctx *gin.Context) {
	var inputData atom_masterdata_food.PutDeleteFoodRequest

	if err := ctx.ShouldBindJSON(&inputData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Invalid request body: " + err.Error(),
		})
		return
	}

	status, err := atom_masterdata_food.PutDeleteFoodUseCase(inputData)

	if !status {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Status error from use case",
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  400,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Successfully deleted food data",
	})
}
