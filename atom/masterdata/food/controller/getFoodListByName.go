package atom_masterdata_food

import (
	"net/http"
	atom_masterdata_food "optimach_service/atom/masterdata/food"

	"github.com/gin-gonic/gin"
)

func GetFoodListByName(ctx *gin.Context) {
	query := ctx.Query("query")
	if query == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Missing 'query' parameter",
			"results": nil,
		})
		return
	}

	data, status, err := atom_masterdata_food.GetFoodListByNameUseCase(query)

	if !status {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
			"results": nil,
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  400,
			"message": err.Error(),
			"results": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Successfully get food data by name",
		"results": gin.H{
			"foods": gin.H{
				"food": data,
			},
		},
	})
}
