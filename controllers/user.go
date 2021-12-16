package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/blakeroy01/internet-orders/models"
	"github.com/blakeroy01/internet-orders/mysql"
)

// Createuser will receive the users information as a request
//  and then store the Users information in the user model in the MySQL Database Server
func CreateUser(requestBody *io.ReadCloser, db *mysql.MySQLDatabase) {
	user := models.User{}
	json.NewDecoder(*requestBody).Decode(&user)

	user.Username = "'JebandBlake'"
	user.Email = "'email@email.com'"
	user.Password = "'password123'"
	user.Active = 1

	sql := fmt.Sprintf("INSERT INTO User(username,email,password,active) VALUES (%s,%s,%s,%d)", user.Username, user.Email, user.Password, user.Active)
	_, err := db.Connection.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}

// Read will receive the user model contents of a request
// then retreive the specified user model in the MySQL Database Server
func ReadUser(requestBody *io.ReadCloser, db *mysql.MySQLDatabase) {
	user := models.User{}
	json.NewDecoder(*requestBody).Decode(&user)
	user.Email = "email@email.com"
	sql := fmt.Sprintf("SELECT email, password FROM user WHERE email = %s ", user.Email)
	results := db.Connection.QueryRow(sql)
	log.Println(results)
	//pass in username and password
	//Find user in the database
}

// Update will receive the body contents of a request
// then update the specified item in the MySQL Database Server
func UpdateUser(requestBody *io.ReadCloser, db *mysql.MySQLDatabase) {
	user := models.User{}
	user.Email = "email@email.com"
	user.Password = "'password123'"
	var newemail = "newemail@newemail.com"
	json.NewDecoder(*requestBody).Decode(&user)
	sql := fmt.Sprintf("UPDATE User SET email = %s WHERE password = %s ", newemail, user.Password)
	results := db.Connection.QueryRow(sql)
	log.Println(results)
}

// Delete will receive the body contents of a request
// then delete the specified item in the MySQL Database Server
func DeleteUser(requestBody *io.ReadCloser, db *mysql.MySQLDatabase) {
	user := models.User{}
	json.NewDecoder(*requestBody).Decode(&user)

	sql := fmt.Sprintf("DELETE FROM user WHERE user_name = %s, password = %s;", user.Username, user.Password)

	_, err := db.Connection.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}
