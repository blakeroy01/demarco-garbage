package main

import (
	"log"
	"net/http"

	"github.com/blakeroy01/internet-orders/mysql"
	"github.com/blakeroy01/internet-orders/transport"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Server Setup
	db, err := mysql.Connect()
	if err != nil {
		log.Fatal(err)
	}

	APIHandler := transport.APIHandler{
		DB: db,
	}

	// Start Server
	log.Println("Starting Server")

	multiplexer := APIHandler.CreateMultiplexer()
	http.ListenAndServe(":4000", multiplexer)

	// Graceful Shutdown
}
