package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetConfig() *AppConfig {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("%v", err)
	}

	conf := &AppConfig{}
	if err := viper.Unmarshal(conf); err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return conf
}
