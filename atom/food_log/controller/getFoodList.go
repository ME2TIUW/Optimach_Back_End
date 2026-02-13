package atom_food_log

import (
	"net/http"
	atom_food_log "optimach_service/atom/food_log"

	"github.com/gin-gonic/gin"
)


func GetFoodList( ctx *gin.Context){

	data, status, err := atom_food_log.GetFoodListUseCase()

	if !status {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status" : 400,
			"message" : err.Error(),
		})

		return 
	}

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status" : 400,
			"message" : err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data" : data,
		"status" : 200,
		"message" : "Succesfully get food list data",
	})
}