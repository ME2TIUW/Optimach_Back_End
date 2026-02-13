package atom_masterdata_food

import (
	"github.com/gin-gonic/gin"
	"net/http"
	atom_masterdata_food "optimach_service/atom/masterdata/food"
)

func GetAllFoodList(ctx *gin.Context) {
	data, status, err := atom_masterdata_food.GetAllFoodListUseCase()

	if !status {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Status error from use case",
		})
		return
	}

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
		"data":    data,
		"status":  200,
		"message": "Successfully get all food list data",
	})
}
