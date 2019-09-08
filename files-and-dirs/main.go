package main

import (
	"bufio"
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	if _, err := os.Stat("file.txt"); os.IsNotExist(err) {
		fmt.Println("file.txt does not exist!")
	}

	contentBytes, err := ioutil.ReadFile("file.txt")
	if err == nil {
		var contentStr = string(contentBytes)
		fmt.Println(contentStr)
	}
	hello := "Hello world"
	err2 := ioutil.WriteFile("helloworld.txt", []byte(hello), 0644)
	if err2 != nil {
		fmt.Println("could not write file")
	}

	// creating temp files
	temFile, err := ioutil.TempFile("", "hello_world_temp")
	if err != nil {
		panic(err)
	}
	defer os.Remove(temFile.Name())

	if _, err := temFile.Write([]byte(hello)); err != nil {
		panic(err)
	}

	fmt.Println(temFile.Name())

	// count lines in file
	file, _ := os.Open("file.txt")
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	defer file.Close()
	fmt.Println(lineCount)

	// read particular line in file
	lineCount = 0
	fileScanner = bufio.NewScanner(file)
	for fileScanner.Scan() {
		fmt.Println("==")
		lineCount++
		if lineCount == 10 {
			fmt.Println(fileScanner.Text())
		}
	}



}
