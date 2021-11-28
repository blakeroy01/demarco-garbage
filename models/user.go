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

// NewUser serves as a constructor for a user object. This object can be used to store input data from the DB or a frontend request
func NewUser(username string, email, string, password string, active int, cardId int, companyId int) *User {
	return &User{
		Username:  username,
		Email:     email,
		Password:  password,
		Active:    active,
		CardID:    cardId,
		CompanyID: companyId,
	}
}
