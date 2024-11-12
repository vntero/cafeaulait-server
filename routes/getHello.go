package routes

import (
	"cafeaulait-server/functions"
	"log"
	"net/http"
)

func GetHello() {
	http.HandleFunc("/", functions.HelloHandler)
	log.Println("GET / is now available")
}
