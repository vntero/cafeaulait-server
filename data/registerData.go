package data

type ResgisterData struct {
	Name                 string `json:"name" validate:"required"`
	Birthday             string `json:"birthday" validate:"required"`
	Origin               string `json:"origin"`
	Motivation           string `json:"motivation"`
	ParentOneName        string `json:"parent_one_name" validate:"required"`
	ParentOneEmail       string `json:"parent_one_email" validate:"required"`
	ParentOnePhone       string `json:"parent_one_phone" validate:"required"`
	ParentOneStreet      string `json:"parent_one_street" validate:"required"`
	ParentOneHouseNumber string `json:"parent_one_house_number" validate:"required"`
	ParentOnePostcode    string `json:"parent_one_postcode" validate:"required"`
	ParentOneLocation    string `json:"parent_one_location" validate:"required"`
	ParentTwoName        string `json:"parent_two_name"`
	ParentTwoEmail       string `json:"parent_two_email"`
	ParentTwoPhone       string `json:"parent_two_phone"`
	ParentTwoStreet      string `json:"parent_two_street"`
	ParentTwoHouseNumber string `json:"parent_two_house_number"`
	ParentTwoPostcode    string `json:"parent_two_postcode"`
	ParentTwoLocation    string `json:"parent_two_location"`
}