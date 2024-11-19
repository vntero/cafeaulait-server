package routes

import (
	"cafeaulait-server/handlers"
	"log"
	"net/http"
)

func PostBook() {
	http.HandleFunc("/book", handlers.AddBook)
	log.Println("POST /book is now available")
}
