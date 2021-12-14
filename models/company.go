package models

// the company model is used to represent companies

type Company struct {
	ID         int64
	Name       string
	Password   string
	BaseCharge int64
}
