package main

import (
	"cafeaulait-server/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// load the env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	routes.GetHello()
	routes.PostBook()
	routes.PostRegister()

	log.Println("Starting server on :" + os.Getenv("PORT"))

	http.ListenAndServe(":"+os.Getenv("PORT"), nil) // start the server
}
