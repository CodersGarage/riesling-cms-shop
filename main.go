package main

import (
	"riesling-cms-core/app/cmd"
	"fmt"
	"os"
	"riesling-cms-core/app/config"
)

func main() {
	config.Init()
	cmd.Init()
	err := cmd.RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
