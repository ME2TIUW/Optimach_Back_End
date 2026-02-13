package atom_activity_log

import (
	"net/http"
	atom_activity_log "optimach_service/atom/activity_log"

	"github.com/gin-gonic/gin"
)

func GetAllActiveActivityLogList (ctx *gin.Context) {

	data, status, err := atom_activity_log.GetAllActiveActivityLogListUseCase()

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
		"message" : "Succesfully get active activity log list data",
	})
}