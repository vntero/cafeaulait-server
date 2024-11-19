package routes

import (
	"cafeaulait-server/handlers"
	"log"
	"net/http"
)

func PostRegister() {
	http.HandleFunc("/register", handlers.AddRegister)
	log.Println("POST /register is now available")
}
