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
	registerData := models.RegisterInput{
		Name:                 r.FormValue("name"),
		Birthday:             r.FormValue("birthday"),
		Origin:               r.FormValue("origin"),
		Motivation:           r.FormValue("motivation"),
		ParentOneName:        r.FormValue("parent_one_name"),
		ParentOneEmail:       r.FormValue("parent_one_email"),
		ParentOnePhone:       r.FormValue("parent_one_phone"),
		ParentOneStreet:      r.FormValue("parent_one_street"),
		ParentOneHouseNumber: r.FormValue("parent_one_house_number"),
		ParentOnePostcode:    r.FormValue("parent_one_postcode"),
		ParentOneLocation:    r.FormValue("parent_one_location"),
		ParentTwoName:        r.FormValue("parent_two_name"),
		ParentTwoEmail:       r.FormValue("parent_two_email"),
		ParentTwoPhone:       r.FormValue("parent_two_phone"),
		ParentTwoStreet:      r.FormValue("parent_two_street"),
		ParentTwoHouseNumber: r.FormValue("parent_two_house_number"),
		ParentTwoPostcode:    r.FormValue("parent_two_postcode"),
		ParentTwoLocation:    r.FormValue("parent_two_location"),
	}

	// print the data to the console
	log.Println("Received Register form submission:", registerData)
}
