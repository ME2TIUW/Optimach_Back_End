package auth_calorie_diary

import (
	"net/http"
	atom_calorie_diary "optimach_service/atom/calorie_diary"

	"github.com/gin-gonic/gin"
)

func GetTotalCalorieDate(ctx *gin.Context) {
	var inputData atom_calorie_diary.GetTotalCalorieByDateRequestModel
	if err := ctx.ShouldBindJSON(&inputData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "invalid request body",
			"datas":   nil,
		})
		return
	}

	data, status, err := atom_calorie_diary.GetTotalCalorieByDateUseCase(inputData.Id_User, inputData.Date)

	if !status {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
			"data":    nil,
		})

		return
	}

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  400,
			"message": err.Error(),
			"data":    nil,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    data,
		"status":  200,
		"message": "succesfully get Total calorie data",
	})
}
