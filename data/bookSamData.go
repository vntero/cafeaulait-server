package data

type BookSamData struct {
	Name           string `json:"name" validate:"required"`
	Email          string `json:"email" validate:"required"`
	Phone          string `json:"phone"`
	EventDate      string `json:"event_date"`
	Location       string `json:"location"`
	NumberOfGuests string `json:"number_of_guests"`
	Occasion       string `json:"occasion" validate:"required"`
	Message        string `json:"message"`
}
