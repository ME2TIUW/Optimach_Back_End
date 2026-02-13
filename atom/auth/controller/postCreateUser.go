package atom_auth

import (
	"log"
	"net/http"
	atom_auth "optimach_service/atom/auth"

	"github.com/gin-gonic/gin"
)

func PostCreateUser(ctx *gin.Context) {

	var inputData atom_auth.UserCredentialsRequestModel

	inputError := ctx.ShouldBindJSON(&inputData)

	if inputError != nil {
		log.Println("[atom][auth][controller][PostCreateUser] invalid request body", inputError.Error())

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Invalid Request Body",
		})

		return
	}

	if inputData.Username == "" || inputData.Password == "" {
		log.Println("[atom][auth][controller][PostRegisterUser] missing required fields",
			"user_name", inputData.Username != "", "",
			"user_password", inputData.Password != "", "",
		)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Username and password can't be empty",
		})

		return
	}

	status, err := atom_auth.PostCreateUserUseCase(inputData)

	if !status {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  400,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Succesfull user registration!",
	})

}
