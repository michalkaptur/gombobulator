package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func hello_handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "I'm the mighty calculator, the result is 42\n")
}

func add_handler(w http.ResponseWriter, req *http.Request) {
	var numbers []int
	err := json.NewDecoder(req.Body).Decode(&numbers)
	if err != nil {
		log.WithError(err).Warn() //todo return HTTP error
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(add(numbers))
}

func add(numbers []int) int {
	//TODO consider int size
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{"remote_addr": r.RemoteAddr, "method": r.Method, "url": r.URL}).Debug("Got a request")
		handler.ServeHTTP(w, r)
	})
}

func main() {
	log.SetLevel(log.DebugLevel) //TODO configure
	http.HandleFunc("/", hello_handler)
	http.HandleFunc("/add", add_handler)
	log.Info("starting the server")
	err := http.ListenAndServe(":7890", logRequest(http.DefaultServeMux))
	if err != nil {
		log.WithError(err).Error("Failed to start the server")
	}
}
