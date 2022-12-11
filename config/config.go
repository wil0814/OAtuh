package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

var Val Config

type Config struct {
	Mode            string `mapstructure:"MODE"`
	Port            string `mapstructure:"PORT"`
	GoogleSecretKey string `mapstructure:"GOOGLE_SECRET_KEY"`
	GoogleClientID  string `mapstructure:"GOOLE_CLIENT_ID"`
	JWTTokenLife    int    `mapstructure:"JWT_TOKEN_LIFE"`
	JWTSecret       string `mapstructure:"JWT_SECRET"`
}

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("讀取設定黨出現錯誤: %v", err))
	}
	if err := viper.Unmarshal(&Val); err != nil {
		panic(fmt.Errorf("找不到Struct, %v", err))
	}
	log.WithFields(log.Fields{
		"val": Val,
	}).Info("config loaded")
}
