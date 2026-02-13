package atom_food_log

import (
	"log"
	"net/http"
	atom_food_log "optimach_service/atom/food_log"

	"github.com/gin-gonic/gin"
)

func GetDEtailFoodLogListByIdUser(ctx *gin.Context) {

	var inputData atom_food_log.DetailFoodLogRequestList

	inputError := ctx.ShouldBindJSON(&inputData)

	if inputError != nil {

		log.Println("[atom][food_log][controller][PutUpdateFoodList] invalid request body", inputError.Error())

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Invalid Request body",
			"error":   inputError.Error(),
		})

		return
	}

	data, status, err := atom_food_log.GetDetailFoodLogByIdUserUseCase(inputData.Id_User, inputData.Created_Date, inputData.Timezone)

	if !status {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
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
		"message": "Succesfully get detail food log list by id user",
	})

}
