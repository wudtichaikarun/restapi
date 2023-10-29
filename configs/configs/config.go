package configs

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Configuration struct {
	Server          ServerConfig
	ExternalService ExternalServiceConfig
}

func GetConfig() (*Configuration, error) {
	var configuration Configuration

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("error to decode, %v", err)
		return nil, err
	}

	isDebug := configuration.Server.Debug

	if isDebug {
		fmt.Println("start using DEV env")
	} else {
		fmt.Println("start using PROD env")
	}

	return &configuration, nil
}
