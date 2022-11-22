package routs

import (
	"encoding/json"
	"net/http"

	"github.com/wilyleo/apiPostgreSql/database"
	"github.com/wilyleo/apiPostgreSql/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get users"))
}
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user"))
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.Profile
	json.NewDecoder(r.Body).Decode(&user)

	createUser := database.DB.Create(&user)
	err := createUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //error 400
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&user)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}
