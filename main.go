package main

import (
	"net/http"

	"github.com/wilyleo/apiPostgreSql/routs"

	"github.com/gorilla/mux"
	"github.com/wilyleo/apiPostgreSql/database"
	"github.com/wilyleo/apiPostgreSql/migration"
)

func main() {

	database.DBConnection()
	migration.Migration()
	r := mux.NewRouter()
	r.HandleFunc("/", routs.Homehandler)
	r.HandleFunc("/user", routs.GetUsersHandler).Methods("GET")
	r.HandleFunc("/user/{id}", routs.GetUserHandler).Methods("GET")
	r.HandleFunc("/user", routs.CreateUser).Methods("POST")
	r.HandleFunc("/user", routs.DeleteUser).Methods("DELETE")
	http.ListenAndServe(":3000", r)

}
