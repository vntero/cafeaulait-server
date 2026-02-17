package main

import (
	"cafeaulait-server/configs"
	"cafeaulait-server/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	// load the env file
	configs.LoadEnv()

	// load routes
	routes.GetHello()
	routes.PostBook()
	routes.PostRegister()
	routes.PostBookSam()

	// start server
	log.Println("Starting server on :" + os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), nil) // start the server
}
