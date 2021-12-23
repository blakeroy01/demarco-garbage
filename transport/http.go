package transport

import (
	"net/http"

	"github.com/blakeroy01/internet-orders/controllers"
	"github.com/blakeroy01/internet-orders/mysql"
)

type APIHandler struct {
	DB *mysql.MySQLDatabase
}

func (handler *APIHandler) HomePage() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Server Started"))
	})
}

func (handler *APIHandler) PaymentInfo() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			controllers.CreateCard(&req.Body, handler.DB)
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
			controllers.CreateUser(&req.Body, handler.DB)
		} else {
			http.Error(res, "Unsupported HTTP Method", http.StatusBadRequest)
		}
	})
}

func (handler *APIHandler) LoginUser() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
			res.Write([]byte("GET"))
		} else if req.Method == http.MethodPost {
			controllers.ReadUser(&req.Body, handler.DB)
		} else {
			http.Error(res, "Unsupported HTTP Method", http.StatusBadRequest)
		}
	})
}

func (handler *APIHandler) UpdateUser() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			controllers.UpdateUser(&req.Body, handler.DB)
		} else {
			http.Error(res, "Unsupported HTTP Method", http.StatusBadRequest)
		}
	})
}

func (handler *APIHandler) DeleteUser() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodDelete {
			controllers.DeleteUser(&req.Body, handler.DB)
		} else {
			http.Error(res, "Unsupported HTTP Method", http.StatusBadRequest)
		}
	})
}

func (handler *APIHandler) CreateMultiplexer() *http.ServeMux {
	multiplexer := http.NewServeMux()
	multiplexer.Handle("/", handler.HomePage())
	multiplexer.Handle("/user/register", handler.RegisterUser())
	multiplexer.Handle("/user/login", handler.LoginUser())
	multiplexer.Handle("/user/delete", handler.DeleteUser())
	multiplexer.Handle("/user/update", handler.UpdateUser())
	multiplexer.Handle("/payments", handler.PaymentInfo())

	return multiplexer
}
