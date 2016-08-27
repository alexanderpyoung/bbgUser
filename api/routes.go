package api

import "github.com/gorilla/mux"

func Handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", createUserHandler).Methods("POST")
	r.HandleFunc("/users", viewAccountHandler).Methods("GET")
	return r
}
