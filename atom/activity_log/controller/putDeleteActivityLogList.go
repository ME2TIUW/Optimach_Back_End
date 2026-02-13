package atom_activity_log

import (
	"log"
	"net/http"
	atom_activity_log "optimach_service/atom/activity_log"

	"github.com/gin-gonic/gin"
)

func PutDeleteActivityLog(ctx *gin.Context) {

	var idActivities atom_activity_log.DeleteActivityLogRequestModel

	inputError := ctx.ShouldBindJSON(&idActivities)

	if inputError != nil {

		log.Println("[atom][food_log][controller][PutUpdateFoodList] invalid request body", inputError.Error())

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status" : 400,
			"message" : "Invalid Request body",
			"error" : inputError.Error(),
		})

		return
	}

	status, err := atom_activity_log.PutDeleteActivityLogUseCase(idActivities.Id_Activities)

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
		"message" : "Succesfully delete activity log list data",
	})
	
}