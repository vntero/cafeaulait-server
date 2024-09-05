package main

import (
	"log"
	"net/http"
)




func main() {
	http.HandleFunc("/", ) // register the route and handler

	log.Println("Starting server on :8080")

	http.ListenAndServe(":8080", nil) // start the server on port 8080
}