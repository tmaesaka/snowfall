package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

var config struct {
	numWorkersPerRequest int
	port                 int
}

var services = [...]string{
	"http://localhost:8081",
	"http://localhost:8081",
	"http://localhost:8081",
	"http://localhost:8081",
}

func initConfig() {
	config.numWorkersPerRequest = 4
	config.port = 8080
}

func spawnWorkers(jobs <-chan string, results chan<- string) {
	for i := 0; i < config.numWorkersPerRequest; i++ {
		go worker(i, jobs, results)
		fmt.Fprintf(os.Stderr, "> Spawned worker: %d\n", i)
	}
}

// TODO: Build a proper JSON response with "encoding/json".
func buildResponse(results []string, elapsed time.Duration) string {
	rv := ""

	for _, result := range results {
		rv += result + "\n"
	}

	rv += "\n==========================\n"
	rv += fmt.Sprintf("Request served in: %.2fs\n", elapsed.Seconds())
	return rv
}

func worker(id int, jobs <-chan string, results chan<- string) {
	for url := range jobs {
		if id%2 == 0 {
			time.Sleep(time.Second * 2)
		}

		resp, err := http.Get(url)

		if err != nil {
			results <- "request failed"
			return
		}

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		results <- fmt.Sprintf("worker_id: %d, %s", id, string(body))
	}
}

func workerHandler(w http.ResponseWriter, r *http.Request) {
	jobs := make(chan string, config.numWorkersPerRequest)
	results := make(chan string, config.numWorkersPerRequest)

	spawnWorkers(jobs, results)
	start := time.Now()

	for _, url := range services {
		jobs <- url
	}

	buf := make([]string, len(services))
	for i := 0; i < len(services); i++ {
		buf[i] = <-results
	}

	res := buildResponse(buf, time.Since(start))
	fmt.Fprintf(w, res)
}

func main() {
	initConfig()

	fmt.Fprintf(os.Stderr, "Starting frontend on port %d...\n", config.port)

	http.HandleFunc("/", workerHandler)
	http.ListenAndServe(":"+strconv.Itoa(config.port), nil)
}
