package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	fmt.Println("copying file")
	file, err := os.Open("original.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file_copy, err2 := os.Create("copy.txt")
	if err2 != nil {
		panic(err2)
	}

	defer file_copy.Close()

	_, err3 := io.Copy(file_copy, file)

	if err3 != nil {
		panic(err3)
	}
}
