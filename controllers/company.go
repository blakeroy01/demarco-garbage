package controllers

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/blakeroy01/internet-orders/models"
	"github.com/blakeroy01/internet-orders/mysql"
)

// Create will receive the body contents of a request
// then store the item in the MySQL Database Server
func CreateCompany(requestBody *io.ReadCloser, db *mysql.MySQLDatabase) error {
	company := models.Company{}
	json.NewDecoder(*requestBody).Decode(&company)

	sql := fmt.Sprintf("INSERT INTO company(name, base_charge, password) VALUES (%s, %v, %s);", company.Name, company.BaseCharge, company.Password)

	_, err := db.Connection.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

// Read will receive the body contents of a request
// then retreive the specified item in the MySQL Database Server
func ReadCompany(requestBody *io.ReadCloser, db *mysql.MySQLDatabase) (*models.Company, error) {
	company := models.Company{}
	json.NewDecoder(*requestBody).Decode(&company)

	sql := fmt.Sprintf("SELECT * FROM card WHERE user_id = %v;", company.ID)

	row := db.Connection.QueryRow(sql)
	var companyData *models.Company

	err := row.Scan(companyData)
	if err != nil {
		return companyData, err
	}

	return companyData, nil
}

// Update will receive the body contents of a request
// then update the specified item in the MySQL Database Server
func UpdateCompany(requestBody *io.ReadCloser, db *mysql.MySQLDatabase) error {
	company := models.Company{}
	json.NewDecoder(*requestBody).Decode(&company)

	sql := fmt.Sprintf("UPDATE company SET name = %s, base_charge = %s WHERE id = %v;", company.Name, company.BaseCharge, company.ID)

	_, err := db.Connection.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

// Delete will receive the body contents of a request
// then delete the specified item in the MySQL Database Server
func DeleteCompany(requestBody *io.ReadCloser, db *mysql.MySQLDatabase) error {
	company := models.Company{}
	json.NewDecoder(*requestBody).Decode(&company)

	sql := fmt.Sprintf("DELETE FROM company WHERE name = %s;", company.Name)

	_, err := db.Connection.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}
