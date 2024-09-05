package functions

import (
	"cafeaulait-server/models"
	"log"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// parse incoming data
	if err := r.ParseForm(); err != nil {
		log.Println("Error parsing data:", err)
		http.Error(w, "Unable to process data", http.StatusBadRequest)
		return
	}

	// create an instance of the struct
	registerData := models.RegisterInput {
		Name: r.FormValue("name"),
		Birthday: r.FormValue("birthday"),
		Origin: r.FormValue("origin"),
	}

	// print the data to the console
	log.Println("Received Register form submission:", registerData)
}