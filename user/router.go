package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

func UserRouter(r *mux.Router) {
	uhandler := NewUserHttpHdl()

	ur := r.PathPrefix("/users").Subrouter()
	ur.HandleFunc("", uhandler.ListUsers).Methods(http.MethodGet)
}
