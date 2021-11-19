package transport

import "net/http"

type APIHandler struct {
}

func (handler *APIHandler) HomePage() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Server Started"))
	})
}

// RegisterUser is a handler function that is used to retrieve the register user page
// Create a user in the database
func (handler *APIHandler) RegisterUser() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// What HTTP Method is sent to this endpoint?
		if req.Method == http.MethodGet {
			res.Write([]byte("GET"))
			// Display HTML to requesting user
		} else if req.Method == http.MethodPost {
			res.Write([]byte("POST"))
		} else {
			res.Write([]byte("Error"))
		}
	})
}

// TODO: (Jeb) Shell out a REST endpoint that handles a POST request to create an order. We can implement a database after we meet.

func (handler *APIHandler) CreateMultiplexer() *http.ServeMux {
	multiplexer := http.NewServeMux()
	multiplexer.Handle("/", handler.HomePage())
	multiplexer.Handle("/user/register", handler.RegisterUser())

	return multiplexer
}
