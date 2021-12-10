package models

// The card model is used to represent card info

type Card struct {
	CardNumber     string
	NameOnCard     string
	CVV            string
	ExpirationDate string
	BillingAddress string
	BillingZip     string
	BillingState   string
	UserId         int
}
