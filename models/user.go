package models

// The user model is used to represent a user within the datebase

type User struct {
	Username  string
	Email     string
	Password  string
	Active    int
	CardID    int
	CompanyID int
}
