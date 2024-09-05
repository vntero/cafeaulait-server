package functions

import (
	"cafeaulait-server/models"
	"log"
	"net/http"
)

func BookHandler(w http.ResponseWriter, r *http.Request) {
	// parse incoming data
	if err := r.ParseForm(); err != nil {
		log.Println("Error parsing data:", err)
		http.Error(w, "Unable to process data", http.StatusBadRequest)
		return
	}

	// create an instance of the struct
	bookData := models.BookInput{
		Name:           r.FormValue("name"),
		Phone:          r.FormValue("phone"),
		Email:          r.FormValue("email"),
		Location:       r.FormValue("location"),
		Duration:       r.FormValue("duration"),
		NumberOfGuests: r.FormValue("number_of_guests"),
		EventDate:      r.FormValue("event_date"),
		EventTime:      r.FormValue("event_time"),
		Budget:         r.FormValue("budget"),
		Comment:        r.FormValue("comment"),
	}

	// print the data to the console
	log.Println("Received Book form submission:", bookData)
}
