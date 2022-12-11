package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, payload interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data":   payload,
	})
}
func SystemError(c *gin.Context, errorCode int, payload interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"result":     false,
		"error_code": errorCode,
		"msg":        payload,
	})
}
