package atom_activity_log

import (
	"log"
	"net/http"
	atom_activity_log "optimach_service/atom/activity_log"

	"github.com/gin-gonic/gin"
)

func PostGetDetailActivityLogListByIdUser(ctx * gin.Context) {
	var idUser atom_activity_log.DetailActivityLogResponseModel

	inputError := ctx.ShouldBindJSON(&idUser)

	if inputError != nil {

		log.Println("[atom][activity_log][controller][PutUpdateFoodList] invalid request body", inputError.Error())

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status" : 400,
			"message" : "Invalid Request body",
			"error" : inputError.Error(),
		})

		return
	}

	data, status, err := atom_activity_log.GetDetailActivityLogByIdUserUseCase(idUser.Id_User)

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
		"message" : "Succesfully get detail activity log list by id user",
	})
	

}