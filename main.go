package main

import (
	"log"
	"net/http"
)


type BookShowInput struct {
	Name           string    `json:"name"`
	Phone          int64     `json:"phone"`
	Email          string    `json:"email"`
	Location       string    `json:"location"`
	Duration       int64     `json:"duration"`
	NumberOfGuests int64     `json:"number_of_guests"`
	EventDate      string `json:"event_date"`
	EventTime      string `json:"event_time"`
	Budget         int64     `json:"budget"`
	Comment        string    `json:"comment"`
}

type RegisterInput struct {
	Name                 string    `json:"name"`
	Birthday             string `json:"birthday"`
	Origin               string    `json:"origin"`
	Motivation           string    `json:"motivation"`
	ParentOneName        string    `json:"parent_one_name"`
	ParentOneEmail       string    `json:"parent_one_email"`
	ParentOnePhone       int64     `json:"parent_one_phone"`
	ParentOneStreet      string    `json:"parent_one_street"`
	ParentOneHouseNumber int64     `json:"parent_one_house_number"`
	ParentOnePostcode    string    `json:"parent_one_postcode"`
	ParentOneLocation    string    `json:"parent_one_location"`
	ParentTwoName        string    `json:"parent_two_name"`
	ParentTwoEmail       string    `json:"parent_two_email"`
	ParentTwoPhone       int64     `json:"parent_two_phone"`
	ParentTwoStreet      string    `json:"parent_two_street"`
	ParentTwoHouseNumber int64     `json:"parent_two_house_number"`
	ParentTwoPostcode    string    `json:"parent_two_postcode"`
	ParentTwoLocation    string    `json:"parent_two_location"`
}



func homeHandler(w http.ResponseWriter, r *http.Request) {
	// parse incoming data
	if err := r.ParseForm(); err != nil {
		log.Println("Error parsing data:", err)
		http.Error(w, "Unable to process data", http.StatusBadRequest)
		return
	}

	registerData := RegisterInput {
		Name: r.FormValue("name"),
		Birthday: r.FormValue("birthday"),
		Origin: r.FormValue("origin"),
	}
}

func main() {
	http.HandleFunc("/", homeHandler) // register the route and handler

	log.Println("Starting server on :8080")

	http.ListenAndServe(":8080", nil) // start the server on port 8080
}