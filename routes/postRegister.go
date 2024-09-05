package routes

import (
	"cafeaulait-server/functions"
	"log"
	"net/http"
)

func PostRegister() {
	http.HandleFunc("/register", functions.RegisterHandler)
	log.Println("POST /register is now available")
}
