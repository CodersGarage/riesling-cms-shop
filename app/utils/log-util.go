package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var logFile, err = os.Create(viper.GetString("others.log_file"))

func LogD(tag string, msg interface{}) {
	if viper.GetString("app.mode") == "debug" {
		fmt.Print(tag)
		fmt.Print(" : ")
		fmt.Println(msg)
	}
}

func LogF(tag string, msg interface{}) {
	if err != nil {
		logFile.WriteString(fmt.Sprintf("%s : %s", tag, msg))
	}
}

func LogP(tag string, msg interface{}) {
	LogD(tag, msg)
	LogF(tag, msg)
}
