package main

import (
	"fmt"
	"os"
	"riesling-cms-shop/app/config"
	"riesling-cms-shop/app/cmd"
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
