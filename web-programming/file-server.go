package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./images")))
	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		panic(err)
	}

}
