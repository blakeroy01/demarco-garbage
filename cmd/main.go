package main

import (
	"net/http"

	"github.com/blakeroy01/internet-orders/transport"
)

func main() {
	// Server Setup
	APIHandler := transport.APIHandler{}
	// Start Server
	multiplexer := APIHandler.CreateMultiplexer()
	http.ListenAndServe(":3000", multiplexer)
	// Graceful Shutdown
	// TODO
}
