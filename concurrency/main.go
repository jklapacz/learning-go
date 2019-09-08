package main

import (
	"fmt"
	"time"
)

func main() {
	nameChannel := make(chan string)
	go func() {
		names := []string{"kuba", "emily", "filip"}
		for _, name := range names {
			time.Sleep(1 * time.Second)
			nameChannel <- name
		}
		close(nameChannel)
	}()

	for data := range nameChannel {
		fmt.Println(data)

	}
	//
	//	go func() {
	//		ages := []int{24, 23, 19}
	//		for _, age := range ages {
	//			time.Sleep(1 * time.Second)
	//			fmt.Println(age)
	//		}
	//	}()
	//
	//	time.Sleep(10 * time.Second)

}
