package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	one, err := ioutil.ReadFile("one.txt")
	if err != nil {
		panic(err)
	}
	two, err := ioutil.ReadFile("two.txt")
	if err != nil {
		panic(err)
	}

	same := bytes.Equal(one, two)
	fmt.Println(same)
}
