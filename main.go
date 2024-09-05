package main

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
type BookShowInput struct {
	Name           string    `json:"name"`
	Phone          int64     `json:"phone"`
	Email          string    `json:"email"`
	Location       string    `json:"location"`
	Duration       int64     `json:"duration"`
	NumberOfGuests int64     `json:"number_of_guests"`
	EventDate      time.Time `json:"event_date"`
	EventTime      time.Time `json:"event_time"`
	Budget         int64     `json:"budget"`
	Comment        string    `json:"comment"`
}

type RegisterInput struct {
	Name                 string    `json:"name"`
	Birthday             time.Time `json:"birthday"`
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



