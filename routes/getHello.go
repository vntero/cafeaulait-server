package routes

import (
	"cafeaulait-server/handlers"
	"log"
	"net/http"
)

func GetHello() {
	http.HandleFunc("/", handlers.ListHello)
	log.Println("GET / is now available")
}
