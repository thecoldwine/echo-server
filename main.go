package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received:", r.Method, r.ContentLength, r.URL)
		r.Write(w)
	})

	bindAddress := os.Getenv("ECHO_SERVER_BIND_ADDRESS")
	if strings.TrimSpace(bindAddress) == "" {
		bindAddress = ":8080"
	}

	log.Println("Starting echo server on", bindAddress)

	err := http.ListenAndServe(bindAddress, mux)
	if err != nil {
		panic(err)
	}
}
