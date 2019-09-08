package main

import (
	"log"
	"os"
	"fmt"
)

type MyError struct {
	ShortMessage string
	DetailedMessage string
}

func (e *MyError) Error() string {
	return e.ShortMessage + "\n" + e.DetailedMessage
}

func main() {
	log_file, err := os.Create("log_file")
	if err != nil {
		fmt.Println("an error occurred...")
	}
	defer log_file.Close()
	log.SetOutput(log_file)

	_,err2 := doSomething()
	if (err2 != nil) {
		log.Println(err2)
		log.Fatalln(err2)
		log.Println("dont log me")
	}
}

func doSomething() (string, error) {
	return "", &MyError{ShortMessage: "I am short", DetailedMessage: "I am long"}
}
