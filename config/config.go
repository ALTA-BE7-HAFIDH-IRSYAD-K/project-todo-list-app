package config

import (
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"sync"
)

type AppConfig struct {
	Port      int    `yaml:"port"`
	SecretJWT string `yaml:"secretjwt"`
	Database  struct {
		Driver   string `yaml:"driver"`
		Name     string `yaml:"name"`
		Address  string `yaml:"address"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.Port = 8000
	defaultConfig.SecretJWT = "S3CR3T"
	defaultConfig.Database.Driver = "mysql"
	defaultConfig.Database.Name = "todolist"
	defaultConfig.Database.Address = "todolist.cdtrgmjwgfbk.us-east-1.rds.amazonaws.com"
	defaultConfig.Database.Port = 3306
	defaultConfig.Database.Username = "admin"
	defaultConfig.Database.Password = "admin123"

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")
	if err := viper.ReadInConfig(); err != nil {
		log.Info("failed to open file")
		return &defaultConfig
	}

	var finalConfig AppConfig
	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Info("failed to extract external config, use default value")
		return &defaultConfig
	}
	return &finalConfig
}
