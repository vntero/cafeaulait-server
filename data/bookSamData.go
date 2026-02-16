package data

type BookSamData struct {
	Name           string `json:"name" validate:"required"`
	LastName       string `json:"last_name" validate:"required"`
	Email          string `json:"email" validate:"required"`
	Phone          string `json:"phone"`
	Location       string `json:"location"`
	Occasion       string `json:"occasion"`
	NumberOfGuests string `json:"number_of_guests"`
	EventDate      string `json:"event_date"`
	Comment        string `json:"comment"`
}
