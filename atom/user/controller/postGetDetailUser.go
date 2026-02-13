package atom_user

import (
	"log"
	"net/http"
	atom_user "optimach_service/atom/user"

	"github.com/gin-gonic/gin"
)

func PostGetDetailUser(ctx *gin.Context) {
	var inputData atom_user.GetDetailUserRequestList

	inputError := ctx.ShouldBindJSON(&inputData)

	if inputError != nil {
		log.Println("[atom][user][controller][PostGetDetailUser] error on binding JSON", inputError.Error())

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Invalid request body",
		})

		return
	}

	data, status, err := atom_user.GetDetailUserUseCase(inputData.Id_User)

	if !status {
		log.Println("[atom][user][controller][PostGetDetailUser] status error - status = ", status)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
		})
		return
	}

	if err != nil {
		log.Println("[atom][user][controller][PostGetDetailUser] error = ", err)
		ctx.JSON(http.StatusOK, gin.H{
			"status":  400,
			"message": err.Error(),
		})
		return

	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    data,
		"status":  200,
		"message": "succesfully get the detail",
	})

}
