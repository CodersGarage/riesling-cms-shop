package app

import "fmt"

func PrintLog(tag string, msg string) {
	fmt.Print(tag)
	fmt.Print(" : ")
	fmt.Println(msg)
}
