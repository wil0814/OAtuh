package handler

import (
	"go-OAuth/pkg/res"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	res.Success(c, gin.H{
		"msg": "pong",
	})
}
