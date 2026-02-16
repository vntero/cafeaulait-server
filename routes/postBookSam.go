package routes

import (
	"cafeaulait-server/handlers"
	"log"
	"net/http"
)

func PostBookSam() {
	http.HandleFunc("/booksam", handlers.AddBookSam)
	log.Println("POST /booksam is now available")
}
