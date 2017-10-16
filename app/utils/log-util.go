package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var logFile *os.File
var err error

func LogD(tag string, msg interface{}) {
	if viper.GetString("app.mode") == "debug" {
		fmt.Print(tag)
		fmt.Print(" : ")
		fmt.Println(msg)
	}
}

func LogF(tag string, msg interface{}) {
	_, err := os.Stat(viper.GetString("others.log_file"))
	if os.IsNotExist(err) {
		logFile, err = os.Create(viper.GetString("others.log_file"))
	}
	if logFile == nil {
		logFile, err = os.Open(viper.GetString("others.log_file"))
	}
	if err != nil {
		logFile.WriteString(fmt.Sprintf("%s : %v", tag, msg))
	}
}

func LogP(tag string, msg interface{}) {
	LogD(tag, msg)
	LogF(tag, msg)
}
