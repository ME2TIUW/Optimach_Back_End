package atom_handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// This handler verifies the application process is running (leapcell).
func SimpleHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "optimach_service_running",
	})
}
