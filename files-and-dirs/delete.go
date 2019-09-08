package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Println("removing file")
	err := os.Remove("new.txt")
	if err != nil {
		panic(err)
	}

}
