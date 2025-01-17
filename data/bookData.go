package data

type BookData struct {
	Name           string `json:"name" validate:"required"`
	Phone          string `json:"phone" validate:"required"`
	Email          string `json:"email" validate:"required"`
	Organization   string `json:"organization"`
	Location       string `json:"location"`
	Duration       string `json:"duration"`
	NumberOfGuests string `json:"number_of_guests"`
	EventDate      string `json:"event_date"`
	EventTime      string `json:"event_time"`
	Budget         string `json:"budget"`
	Comment        string `json:"comment"`
}
