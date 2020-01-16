package utils

import (
	"log"

	model "github.com/skill-central-golang/models"
	"github.com/spf13/viper"
)

func LoadConfig() *model.Config {
	config := &model.Config{}
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("Fatal error config file: %s \n", err)
	}
	viper.Unmarshal(config)
	return config
}
