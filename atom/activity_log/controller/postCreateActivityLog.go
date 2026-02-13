package atom_activity_log

import (
	"log"
	"net/http"
	atom_activity_log "optimach_service/atom/activity_log"

	"github.com/gin-gonic/gin"
)

func PostCreateActivityLog(ctx *gin.Context) {

	var inputData atom_activity_log.CreateActivityLogListRequestModel

	inputError := ctx.ShouldBindJSON(&inputData)

	if inputError != nil {
		log.Println("[atom][activity_log][controller][PostCreateActivityLog] invalid request body", inputError.Error())

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status" : 400,
			"message" : "Invalid Request body",
			"error" : inputError.Error(),
		})
		return
	}

	status, err := atom_activity_log.PostCreateActivityLogUseCase(inputData)

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
		"message" : "Succesfully create activity log list data",
	})


}