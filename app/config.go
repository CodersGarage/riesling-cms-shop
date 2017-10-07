package app

import (
	"io/ioutil"
	"github.com/tidwall/gjson"
)

var config gjson.Result

func InitAppConfig() {
	value, err := ioutil.ReadFile("./app.json")

	if err != nil {
		panic(err)
	}
	config = gjson.Parse(string(value))
}

func GetAppConfig() gjson.Result {
	return config
}
