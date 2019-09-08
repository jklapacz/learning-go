package main

import (
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	planet := r.URL.Query().Get("planet")
	w.Write([]byte("Hello, " + planet))
}
func sayMars(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, mars"))
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/mars", sayMars)
	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		panic(err)
	}

}
