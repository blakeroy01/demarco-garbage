package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/blakeroy01/internet-orders/models"
	"github.com/blakeroy01/internet-orders/mysql"
)

// Create will receive the body contents of a request
// then store the item in the MySQL Database Server
func CreateCard(requestBody *io.ReadCloser, db *mysql.MySQLDatabase) {
	card := models.Card{}
	json.NewDecoder(*requestBody).Decode(&card)

	// Remove this upon deployment 
	card.CardNumber = "'4444444444444444'"
	card.NameOnCard = "'Jeb & Blakes Credit Card'"
	card.CVV = "'679'"
	card.ExpirationDate = "'02/25'"
	card.BillingAddress = "'123 Fetty Wap Drive'"
	card.BillingZip = "'33409'"
	card.BillingState = "'NY'"

	sql := fmt.Sprintf("INSERT INTO card(card_number, name_on_card, cvv, expiration_date, billing_address, billing_zip, billing_state) VALUES (%s, %s, %s, %s, %s, %s, %s);", card.CardNumber, card.NameOnCard,
		card.CVV, card.ExpirationDate, card.BillingAddress, card.BillingZip,
		card.BillingState)

	_, err := db.Connection.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}

// Read will receive the body contents of a request
// then retreive the specified item in the MySQL Database Server
func ReadCard() {

}

// Update will receive the body contents of a request
// then update the specified item in the MySQL Database Server
func UpdateCard() {

}

// Delete will receive the body contents of a request
// then delete the specified item in the MySQL Database Server
func DeleteCard() {

}
