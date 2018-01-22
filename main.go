package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	server := http.Server{
		Addr:         ":4000",
		WriteTimeout: 10 * time.Second,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	server.Handler = mux
	log.Println("server is running at 4000")
	err := server.ListenAndServe()
	log.Fatal(err)
}

func home(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "hello world")
}
