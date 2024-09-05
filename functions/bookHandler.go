package functions

import (
	"cafeaulait-server/models"
	"fmt"
	"log"
	"net/http"
)

func bookHandler(w http.ResponseWriter, r *http.Request) {
	// parse incoming data
	if err := r.ParseForm(); err != nil {
		log.Println("Error parsing data:", err)
		http.Error(w, "Unable to process data", http.StatusBadRequest)
		return
	}

	// create an instance of the struct
	bookData := models.BookInput {
		Name: r.FormValue("name"),
	}

	// print the data to the console
	log.Println("Received form submission:", bookData)

	fmt.Fprintf(w, "Form submitted successfully!")
}