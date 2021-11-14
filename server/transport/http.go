package transport

import "net/http"

type APIHandler struct {
}

func (handler *APIHandler) HomePage() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Server Started"))
	})
}

// TODO: (Jeb) Shell out a REST endpoint that handles a POST request to create an order. We can implement a database after we meet.

func (handler *APIHandler) CreateMultiplexer() *http.ServeMux {
	multiplexer := http.NewServeMux()
	multiplexer.Handle("/", handler.HomePage())

	return multiplexer
}
