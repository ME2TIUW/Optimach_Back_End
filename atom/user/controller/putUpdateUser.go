package atom_user

import (
	"log"
	"net/http"
	atom_user "optimach_service/atom/user"

	"github.com/gin-gonic/gin"
)

func PutUpdateUser(ctx *gin.Context) {

	var inputData atom_user.UpdateUserRequestList

	inputError := ctx.ShouldBindJSON(&inputData)

	if inputError != nil {
		log.Println("[atom][user][controller][putUpdateUser] invalid request body", inputError.Error())

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"message" : "Invalid request body",
		})
		return 
	}

	status, err := atom_user.PutUpdateUserUseCase(inputData)

	if !status {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"message" : err.Error(),
		})

		return 
	}
	
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": 400,
			"message" : err.Error(),
		})
	
		return 
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status" : 200,
		"message" : "Succesfully update user's data",
	})
}