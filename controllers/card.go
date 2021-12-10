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
func CreateCard(requestBody *io.ReadCloser, db *mysql.MySQLDatabase) error {
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
	card.UserId = 1

	sql := fmt.Sprintf("INSERT INTO card(card_number, name_on_card, cvv, expiration_date, billing_address, billing_zip, billing_state, user_id) VALUES (%s, %s, %s, %s, %s, %s, %s, %v);", card.CardNumber, card.NameOnCard,
		card.CVV, card.ExpirationDate, card.BillingAddress, card.BillingZip,
		card.BillingState, card.UserId)

	_, err := db.Connection.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

// Read will receive the body contents of a request
// then retreive the specified item in the MySQL Database Server
func ReadCard(requestBody *io.ReadCloser, db *mysql.MySQLDatabase) (models.Card, error) {
	card := models.Card{}
	json.NewDecoder(*requestBody).Decode(&card)

	userID := card.UserId

	sql := fmt.Sprintf("SELECT * FROM card WHERE user_id = %v;", userID)

	row := db.Connection.QueryRow(sql)
	var cardData models.Card

	err := row.Scan(&cardData)
	if err != nil {
		return cardData, err
	}

	return cardData, nil

}

// Update will receive the body contents of a request
// then update the specified item in the MySQL Database Server
func UpdateCard(requestBody *io.ReadCloser, db *mysql.MySQLDatabase) {
	card := models.Card{}
	json.NewDecoder(*requestBody).Decode(&card)

	sql := fmt.Sprintf("UPDATE card SET card_number = %s , name_on_card = %s , cvv = %s, expiration_date = %s, billing_address = %s, billing_zip = %s, billing_state = %s) WHERE user_id = %v;", card.CardNumber, card.NameOnCard,
		card.CVV, card.ExpirationDate, card.BillingAddress, card.BillingZip,
		card.BillingState, card.UserId)

	_, err := db.Connection.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}

// Delete will receive the body contents of a request
// then delete the specified item in the MySQL Database Server
func DeleteCard(requestBody *io.ReadCloser, db *mysql.MySQLDatabase) error {
	card := models.Card{}
	json.NewDecoder(*requestBody).Decode(&card)

	sql := fmt.Sprintf("DELETE FROM card WHERE user_id = %v;", card.UserId)

	_, err := db.Connection.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}
