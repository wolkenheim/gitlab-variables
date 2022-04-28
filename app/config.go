package app

import (
	"fmt"
	"github.com/spf13/viper"
)

func ReadConfig(envName string) {
	viper.SetConfigName(envName)
	viper.SetConfigType("yaml")
	if envName == "testing" {
		viper.AddConfigPath("../../config")
	} else {
		viper.AddConfigPath("./config")
	}
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
