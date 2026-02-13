package atom_activity_log

import (
	"log"
	"net/http"
	atom_activity_log "optimach_service/atom/activity_log"

	"github.com/gin-gonic/gin"
)

func PutUpdateActivityLogList (ctx * gin.Context) {

	var inputData atom_activity_log.UpdateActivityLogListRequestModel

	inputError := ctx.ShouldBindJSON(&inputData)

	if inputError != nil {

		log.Println("[atom][activity_log][controller][PutUpdateactivityLogList] invalid request body", inputError.Error())

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status" : 400,
			"message" : "Invalid Request body",
			"error" : inputError.Error(),
		})

		return
	}

	status, err := atom_activity_log.PutUpdateActivityLogUseCase(inputData)

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
		"message" : "Succesfully update activity log list data",
	})


}