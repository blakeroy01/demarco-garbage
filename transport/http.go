package transport

import "net/http"

type APIHandler struct {
}

func (handler *APIHandler) HomePage() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Server Started"))
	})
}

func (handler *APIHandler) PaymentInfo() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// What HTTP Method is sent to this endpoint?
		if req.Method == http.MethodGet {
			res.Write([]byte("GET"))
			// Display HTML to requesting user
		} else if req.Method == http.MethodPost {
			res.Write([]byte("POST"))
		} else if req.Method == http.MethodPut {
			res.Write([]byte("PUT"))
		} else {
			http.Error(res, "Unsupported HTTP Method", http.StatusBadRequest)
		}
	})
}

// RegisterUser is a handler function that is used to create a user in the database
func (handler *APIHandler) RegisterUser() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			res.Write([]byte("POST"))
		} else {
			http.Error(res, "Unsupported HTTP Method", http.StatusBadRequest)
		}
	})
}

func (handler *APIHandler) CreateMultiplexer() *http.ServeMux {
	multiplexer := http.NewServeMux()
	multiplexer.Handle("/", handler.HomePage())
	multiplexer.Handle("/user/register", handler.RegisterUser())
	multiplexer.Handle("/payments", handler.PaymentInfo())

	return multiplexer
}
