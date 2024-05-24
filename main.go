package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "I'm the mighty calculator, the result is 42\n")
}

func add(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "I'm the mighty calculator, but I don't know how to calculate sum just yet\n")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/add", add)
	http.ListenAndServe(":7890", nil)
}
