package utils

import (
	"os"
	"github.com/spf13/viper"
	"fmt"
	"strconv"
	"io/ioutil"
	"strings"
	"github.com/nu7hatch/gouuid"
)

/**
 * := Coded with love by Sakib Sami on 3/11/17.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

func GetAppPID() int {
	return os.Getpid()
}

func ReadPID() string {
	v, err := ioutil.ReadFile(viper.GetString("others.pid_file"))
	if err != nil {
		fmt.Println("Couldn't find PID file")
	}
	return string(v)
}

func PutAppPID() {
	pid := GetAppPID()
	pidFile, err := os.Create(viper.GetString("others.pid_file"))
	if err != nil {
		fmt.Println("Couldn't create PID file")
	}
	pidFile.WriteString(strconv.Itoa(pid))
	pidFile.Close()
}

func GetUUID() string {
	uid, err := uuid.NewV4()
	if err != nil {
		return ""
	}
	return strings.Replace(uid.String(), "-", "", -1)
}
