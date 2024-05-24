package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "I'm the mighty calculator, the result is 42\n")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":7890", nil)
}
