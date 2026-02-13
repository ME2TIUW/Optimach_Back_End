package atom_food_log

import (
	"log"
	"net/http"
	atom_food_log "optimach_service/atom/food_log"

	"github.com/gin-gonic/gin"
)

func PostCreateFoodLog(ctx *gin.Context) {

	var inputData atom_food_log.CreateFoodLogRequestModel

	inputError := ctx.ShouldBindJSON(&inputData)

	if inputError != nil {
		log.Println("[atom][food_log][controller][PostCreateFoodLog] invalid request body", inputError.Error())

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status" : 400,
			"message" : "Invalid Request body",
			"error" : inputError.Error(),
		})

		return
	}

	status, err := atom_food_log.PostCreateFoodLogUseCase(inputData)

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
		"status" : 200,
		"message" : "Succesfully create food list data",
	})


}