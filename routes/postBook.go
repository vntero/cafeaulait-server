package routes

import (
	"cafeaulait-server/functions"
	"log"
	"net/http"
)

func PostBook() {
	http.HandleFunc("/book", functions.BookHandler)
	log.Println("POST /book is now available")
}