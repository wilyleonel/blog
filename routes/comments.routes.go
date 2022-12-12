package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wilyleo/apiPostgreSql/models"
)

func GetAllComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	comment, err := models.GetAllComments()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(&comment)
	}
}
func GetComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	comment, err := models.GetComments(int(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(comment)
}
func CreateComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var comment models.Comments
	err := decoder.Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = models.CreateComments(comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

/*func GetAllCommentsHandler(w http.ResponseWriter, r *http.Request) {
	var allComments []models.Comments
	database.DB.Table("Comments").Find(&allComments)
	json.NewEncoder(w).Encode(&allComments)
}

func GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
	var idComments models.Comments
	params := mux.Vars(r)
	database.DB.First(&idComments, params["id"])

	json.NewEncoder(w).Encode(&idComments)
}
func CreateComments(w http.ResponseWriter, r *http.Request) {
	var newComments models.Comments
	json.NewDecoder(r.Body).Decode(&newComments)

	createComments := database.DB.Create(&newComments)

	err := createComments.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //error 400
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&newComments)
}
*/
