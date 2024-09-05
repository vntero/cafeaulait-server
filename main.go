package main

import (
	"cafeaulait-server/routes"
	"log"
	"net/http"
)

func main() {
	routes.PostBook()
	routes.PostRegister()

	log.Println("Starting server on :8080")

	http.ListenAndServe(":8080", nil) // start the server on port 8080
}
