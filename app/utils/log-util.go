package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var logFile, err = os.Create(viper.GetString("others.log_file"))

func LogD(tag string, msg string) {
	if viper.GetString("app.mode") == "debug" {
		fmt.Print(tag)
		fmt.Print(" : ")
		fmt.Println(msg)
	}
}

func LogF(tag string, msg string) {
	if err != nil {
		logFile.WriteString(fmt.Sprintf("%s : %s", tag, msg))
	}
}
