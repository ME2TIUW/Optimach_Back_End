package atom_masterdata_food

import (
	"net/http"
	atom_masterdata_food "optimach_service/atom/masterdata/food"

	"github.com/gin-gonic/gin"
)

func GetAllActiveFoodList(ctx *gin.Context) {
	data, status, err := atom_masterdata_food.GetAllActiveFoodListUseCase()

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
		"message": "Successfully get active food list data",
	})
}
