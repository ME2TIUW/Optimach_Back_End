package atom_user

import (
	"net/http"
	atom_user "optimach_service/atom/user"

	"github.com/gin-gonic/gin"
)

func PutDeleteUserList (ctx *gin.Context)  {

		var inputData atom_user.DeleteUserRequestList

		inputError := ctx.ShouldBindJSON(&inputData)


		if inputError != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status" : "400",
				"message" : "Invalid request body",
			})
			return
		}

		status, err := atom_user.PutDeleteUserUseCase(inputData)

		if !status {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message" : err.Error(),
				"status" : 400, 
			})
			return
		}
	
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message" : err.Error(),
				"status" : 400, 
			})
			return
		}
		
		ctx.JSON(http.StatusOK, gin.H{
			"message" : "succesfully delete status user list",
			"status" : 200, 
		})

}