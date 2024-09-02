package api

import "time"

type BookShowInput struct {
	Name 			string `json:"name"`
	Phone 			int64 `json:"phone"`
	Email			string `json:"email"`
	Location		string `json:"location"`
	Duration		int64 `json:"duration"`
	NumberOfGuests	int64 `json:"number_of_guests"`
	EventDate		time.Time `json:"event_date"`
	EventTime		time.Time `json:"event_time"`
	Budget			int64 `json:"budget"`
	Comment			string `json:"comment"`
}