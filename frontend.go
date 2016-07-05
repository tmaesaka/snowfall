package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var services = [...]string{
	"http://localhost:8081",
	"http://localhost:8081",
	"http://localhost:8081",
	"http://localhost:8081",
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

func sendRequest(reqId int, url string, ch chan string) {
	if reqId%2 == 0 {
		time.Sleep(time.Second * 2)
	}

	resp, err := http.Get(url)
	if err != nil {
		ch <- "request failed"
		return
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	ch <- fmt.Sprintf("request_id: %d, %s", reqId, string(body))
}

func concurrentRequestsHandler(w http.ResponseWriter, r *http.Request) {
	ch := make(chan string)
	results := make([]string, len(services))

	start := time.Now()

	for i, url := range services {
		go sendRequest(i, url, ch)
	}

	for i := 0; i < cap(results); i++ {
		results[i] = <-ch
	}

	res := buildResponse(results, time.Since(start))
	fmt.Fprintf(w, res)
}

func main() {
	fmt.Fprint(os.Stderr, "Starting frontend on port 8080...\n")

	http.HandleFunc("/", concurrentRequestsHandler)
	http.ListenAndServe(":8080", nil)
}
