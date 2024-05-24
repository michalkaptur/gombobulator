package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "I'm the mighty calculator, the result is 42\n")
}

func add(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "I'm the mighty calculator, but I don't know how to calculate sum just yet\n")
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{"remote_addr": r.RemoteAddr, "method": r.Method, "url": r.URL}).Debug("Got a request")
		handler.ServeHTTP(w, r)
	})
}

func main() {
	log.SetLevel(log.DebugLevel) //TODO configure
	http.HandleFunc("/", hello)
	http.HandleFunc("/add", add)
	log.Info("starting the server")
	err := http.ListenAndServe(":7890", logRequest(http.DefaultServeMux))
	if err != nil {
		log.WithError(err).Error("Failed to start the server")
	}
}
