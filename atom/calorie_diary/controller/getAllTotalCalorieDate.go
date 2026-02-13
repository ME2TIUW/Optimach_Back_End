package auth_calorie_diary

import (
	"net/http"
	atom_calorie_diary "optimach_service/atom/calorie_diary"

	"github.com/gin-gonic/gin"
)

func GetAllTotalCalorieDate(ctx *gin.Context) {

	datas, status, err := atom_calorie_diary.GetAllTotalCalorieDateUseCase()

	if !status {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
			"datas":   nil,
		})

		return
	}

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  400,
			"message": err.Error(),
			"datas":   nil,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"datas":   datas,
		"status":  200,
		"message": "succesfully get Total calorie data",
	})
}
