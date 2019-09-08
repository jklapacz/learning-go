package main

import (
	"fmt"
	"os"
)

func main() {
	realArgs := os.Args[1:]
	if len(realArgs) == 0 {
		fmt.Println("Please pass an argument")
		return
	}
	if realArgs[0] == "a" {
		fmt.Println("Hello world")
	} else if realArgs[0] == "b" {
		fmt.Println("Hello mars")
	} else {
		fmt.Println("unsupported argument")
	}
}
