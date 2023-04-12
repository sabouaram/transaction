package controller

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", CreateUser).Methods("POST").Schemes("http")
	r.HandleFunc("/users/{id:[a-fA-F0-9]{8}-(?:[a-fA-F0-9]{4}-){3}[a-fA-F0-9]{12}}", GetUser).Methods("GET").Schemes("http")
	r.HandleFunc("/users/{id:[a-fA-F0-9]{8}-(?:[a-fA-F0-9]{4}-){3}[a-fA-F0-9]{12}}", UpdateUser).Methods("PUT").Schemes("http")
	r.HandleFunc("/users/{id:[a-fA-F0-9]{8}-(?:[a-fA-F0-9]{4}-){3}[a-fA-F0-9]{12}}", DelUser).Methods("DELETE").Schemes("http")
	r.HandleFunc("/transaction", MakeTransaction).Methods("POST").Schemes("http")
	return r
}
