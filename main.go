package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

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
	sum := add_launcher(numbers)
	json.NewEncoder(w).Encode(sum)
}

func add_launcher(numbers []int) int {
	var wg sync.WaitGroup
	partial_sums := make(chan int)
	if len(numbers)%2 != 0 {
		numbers = append(numbers, 0)
		// pad with additional zero if odd count
	}
	pool_size := len(numbers) / 2
	wg.Add(pool_size)
	for i := range pool_size {
		low := i * 2
		high := ((i + 1) * 2)
		log.WithFields(log.Fields{"low": low, "high": high}).Debug("Launching calculation")
		go add_async(numbers[low:high], partial_sums, &wg)
	}
	sum := 0
	for range pool_size {
		sum += <-partial_sums
	}
	wg.Wait()
	return sum
}

func add_async(numbers []int, output chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	output <- add(numbers)
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
