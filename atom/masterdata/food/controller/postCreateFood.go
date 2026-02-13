package atom_masterdata_food

import (
	"net/http"
	atom_masterdata_food "optimach_service/atom/masterdata/food"

	"github.com/gin-gonic/gin"
)

// PostCreateFood creates a new food item.
func PostCreateFood(ctx *gin.Context) {
	var inputData atom_masterdata_food.PostCreateFoodRequest

	// Bind the request body to the struct
	if err := ctx.ShouldBindJSON(&inputData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Invalid request body: " + err.Error(),
		})
		return
	}

	status, err := atom_masterdata_food.PostCreateFoodUseCase(inputData)

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

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  201,
		"message": "Successfully created new food data",
	})
}
