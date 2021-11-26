package models

// the company model is used to represent companies

type Company struct {
	Name       string
	Password   string
	BaseCharge int64
}

// NewCompany serves as a constructor for a company object. This object can be used to store input data from the DB or a frontend request
func NewCompany(name string, password string, baseCharge int64) *Company {
	return &Company{
		Name:       name,
		Password:   password,
		BaseCharge: baseCharge,
	}
}
