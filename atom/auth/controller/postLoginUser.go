package atom_auth

import (
	"log"
	"net/http"
	atom_auth "optimach_service/atom/auth"

	"github.com/gin-gonic/gin"
)

func PostLoginUser(ctx *gin.Context) {

	var inputData atom_auth.UserCredentialsRequestModel

	inputError := ctx.ShouldBindJSON(&inputData)

	if inputError != nil {
		log.Println("[atom][auth][controller][PostLoginUser] invalid request body", inputError.Error())

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Invalid request body",
		})
		return
	}

	if inputData.Username == "" || inputData.Password == "" {
		log.Println("[atom][auth][controller][PostLoginUser] invalid username/password")

		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"status":  400,
			"message": "input data cannot be empty",
		})
		return

	}

	data, accessToken, refreshToken, _, _, status, err, errPassword := atom_auth.GetAllUserListUseCase(inputData)

	if errPassword != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"status":  400,
			"message": "invalid password",
		})

		return
	}

	if !status {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"status":  400,
			"message": err.Error(),
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"data":    nil,
			"status":  400,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"status":        200,
		"message":       "Successfully logged in",
		"credential": gin.H{
			"id_user":          data.Id_User,
			"username":         data.Username,
			"have_filled_form": data.Have_Filled_Form,
			"is_admin":         data.Is_Admin,
			"is_active":        data.Is_Active,
		},
	})
}
