package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", DefaultHandler)
	http.HandleFunc("/status", StatusHandler)
	fmt.Println("Start Server")
	http.ListenAndServe(":8080", nil)
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "status ok")
}

