package models

type BookInput struct {
	Name           string    `json:"name"`
	Phone          string     `json:"phone"`
	Email          string    `json:"email"`
	Location       string    `json:"location"`
	Duration       string     `json:"duration"`
	NumberOfGuests string     `json:"number_of_guests"`
	EventDate      string 	 `json:"event_date"`
	EventTime      string 	 `json:"event_time"`
	Budget         string     `json:"budget"`
	Comment        string    `json:"comment"`
}

type RegisterInput struct {
	Name                 string    `json:"name"`
	Birthday             string    `json:"birthday"`
	Origin               string    `json:"origin"`
	Motivation           string    `json:"motivation"`
	ParentOneName        string    `json:"parent_one_name"`
	ParentOneEmail       string    `json:"parent_one_email"`
	ParentOnePhone       string     `json:"parent_one_phone"`
	ParentOneStreet      string    `json:"parent_one_street"`
	ParentOneHouseNumber string     `json:"parent_one_house_number"`
	ParentOnePostcode    string    `json:"parent_one_postcode"`
	ParentOneLocation    string    `json:"parent_one_location"`
	ParentTwoName        string    `json:"parent_two_name"`
	ParentTwoEmail       string    `json:"parent_two_email"`
	ParentTwoPhone       string     `json:"parent_two_phone"`
	ParentTwoStreet      string    `json:"parent_two_street"`
	ParentTwoHouseNumber string     `json:"parent_two_house_number"`
	ParentTwoPostcode    string    `json:"parent_two_postcode"`
	ParentTwoLocation    string    `json:"parent_two_location"`
}