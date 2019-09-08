package main

import (
	"fmt"
)

func main() {
	saySomething()
	fmt.Println("After the panicing function")
}

func saySomething() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered in f", r)

		}
	}()
	panicer()
}

func panicer() {
	fmt.Println("==")
	panic("panicking!")
}

