package config

import "github.com/spf13/viper"

func Init() {
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/Users/s4kib/go/src/riesling-cms-shop")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
