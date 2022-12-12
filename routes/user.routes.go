package routes

import (

	// "github.com/gorilla/mux"
	// "github.com/wilyleo/apiPostgreSql/database"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wilyleo/apiPostgreSql/models"
)

const (
	HASH_COST = 8
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := models.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(&users)
	}
}

// allusers
// func GetAllUsers(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	users, err := models.GetAllUsers()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(err.Error()))
// 	} else {
// 		json.NewEncoder(w).Encode(&users)
// 	}
// }

// user/id
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	user, err := models.GetUsers(int(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(user)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var user models.Users
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = models.CreateUser(user, err)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

// func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
// 	var users []models.Users
// 	database.DB.Table("users").Joins("Profile").Joins("Post").Joins("Comments").Find(&users)
// 	json.NewEncoder(w).Encode(&users)

// }
// func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
// 	var users []models.User
// 	database.DB.Table("users").Joins("left join profiles on profiles.user_id = users.id").Find(&users)
// 	json.NewEncoder(w).Encode(&users)
// }

// // user/id
// func GetUserHandler(w http.ResponseWriter, r *http.Request) {
// 	var idUser models.Users
// 	params := mux.Vars(r)
// 	searchUser := database.DB.First(&idUser, params["id"])
// 	err := searchUser.Error
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest) //error 400
// 		w.Write([]byte(err.Error()))
// 	}

// 	json.NewEncoder(w).Encode(&idUser)
// }

// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	var user models.Users
// 	json.NewDecoder(r.Body).Decode(&user)

// 	createUser := database.DB.Create(&user)
// 	err := createUser.Error
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest) //error 400
// 		w.Write([]byte(err.Error()))
// 	}
// 	json.NewEncoder(w).Encode(&user)
// }
// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	var userDelete models.Users
// 	params := mux.Vars(r)
// 	database.DB.First(&userDelete, params["id"])
// 	if userDelete.ID == 0 {
// 		w.WriteHeader(http.StatusNotFound) //error 400
// 		w.Write([]byte("user not found"))
// 	}

// 	database.DB.Delete(&userDelete)
// 	w.WriteHeader(http.StatusOK)
// }
