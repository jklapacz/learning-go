package main

import (
	"fmt"
	"strconv"
)

func main() {
	// converting with package
	isNew := true
	isNewStr := strconv.FormatBool(isNew)
	defMessage := "Purchased item is: "
	message := defMessage + isNewStr
	fmt.Println(message)
	// convert with string formatting
	// %v is a default value for a given variable
	fmt.Printf("%s %v \n", defMessage, isNew)
	// integer + float to string
	number := 100
	numberStr := strconv.Itoa(number)
	fmt.Println(numberStr)
	numberFloat := 123134.12142323434
	numberFloatStr := strconv.FormatFloat(numberFloat, 'f', -1, 64)
	fmt.Println(numberFloatStr)
	// string to boolean
	isNewStr = "1"
	isNewBool, err := strconv.ParseBool(isNewStr)
	// need to handle error
	if (err != nil) {
		fmt.Println("failed")
	} else {
		if (isNewBool) {
			fmt.Println("it is new")
		} else {
			fmt.Println("not new")
		}
	}
	// parse int and float string to int and float types
	numStr := "2"
	valueInt, err := strconv.ParseInt(numStr, 10, 32)
	if (err != nil) {
		fmt.Println("failed")
	} else {
		fmt.Printf("%v \n", valueInt)
	}
	numFlt := "2.01"
	valueFlt, err := strconv.ParseFloat(numFlt, 64)
	if (err != nil) {
		fmt.Println("failed")
	} else {
		fmt.Printf("%v \n", valueFlt)
	}

	// convert byte array to string
	helloWorldByte := []byte {72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100}
	fmt.Println(string(helloWorldByte))
	// convert string to byte array
	nonByte := "Hello world"
	fmt.Println([]byte(nonByte))
}
