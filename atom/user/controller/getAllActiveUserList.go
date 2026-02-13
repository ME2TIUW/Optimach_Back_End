package atom_user

import (
	"net/http"
	atom_user "optimach_service/atom/user"

	"github.com/gin-gonic/gin"
)

func GetAllActiveUserList(ctx *gin.Context) {
	
	datas, status, err := atom_user.GetAllActiveUserListUseCase()
	
	if !status {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data" : nil,
			"message" : err.Error(),
			"status" : 400, 
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"data" : nil,
			"message" : err.Error(),
			"status" : 400, 
		})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"data" : datas,
		"message" : "succesfully see all active user's list",
		"status" : 200, 
	})
}