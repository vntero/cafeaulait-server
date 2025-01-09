package data

type BookData struct {
	Name           string `json:"name"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Location       string `json:"location"`
	Duration       string `json:"duration"`
	NumberOfGuests string `json:"number_of_guests"`
	EventDate      string `json:"event_date"`
	EventTime      string `json:"event_time"`
	Budget         string `json:"budget"`
	Comment        string `json:"comment"`
}