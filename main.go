package main

import (
	"go-OAuth/config"
	"go-OAuth/handler"
	"go-OAuth/middleware"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func main() {
	load()
	gin.SetMode(config.Val.Mode)
	r := gin.Default()
	r.Use(middleware.CROSS())
	r.GET("/ping", handler.Ping)
	api := r.Group("/api")
	{
		api.GET("ouath/google/url", handler.GoogleAccsess)
		api.GET("ouath/google/login", handler.GoogleLogin)
	}
	r.Run(":" + config.Val.Port)
	log.Info("save port: %v \n", config.Val.Port)

}

func load() {
	config.Init()
}
