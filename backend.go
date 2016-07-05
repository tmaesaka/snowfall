package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func timeHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	fmt.Fprintf(w, "%s", time.Now().UTC())
}

func main() {
	fmt.Fprint(os.Stderr, "Starting backend on port 8081...\n")

	http.HandleFunc("/", timeHandler)
	err := http.ListenAndServe(":8081", nil)

	if err != nil {
		log.Fatal("Failed to start backend: ", err)
	}
}
