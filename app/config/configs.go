package config

import "github.com/spf13/viper"

/**
 * := Coded with love by Sakib Sami on 3/11/17.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

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
