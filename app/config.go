package app

import (
	"io/ioutil"
	"github.com/tidwall/gjson"
)

var config gjson.Result

func InitAppConfigWithPath(configPath string) {
	value, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	config = gjson.Parse(string(value))
}

func InitAppConfig() {
	InitAppConfigWithPath("./app.json")
}

func GetAppConfig() gjson.Result {
	return config
}
