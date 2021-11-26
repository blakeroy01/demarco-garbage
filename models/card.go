package models

// The item model is used to represent and create items on a restaurant menu

type Card struct {
	UserID int
	Number int
	Name   string
	CVV    string
}
