package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wilyleo/apiPostgreSql/models"
)

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	categories, err := models.GetAllCategories()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(&categories)
	}
}
func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	category, err := models.GetCategory(int(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&category)
}

func CreateCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var category models.Categories
	err := decoder.Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = models.CreateCategories(category)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

/*func GetAllCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	var allCategories []models.Categories
	database.DB.Table("categories").Find(&allCategories)
	json.NewEncoder(w).Encode(&allCategories)
}
func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	var idCategories models.Comments
	params := mux.Vars(r)
	database.DB.First(&idCategories, params["id"])
	fmt.Println(idCategories)
	json.NewEncoder(w).Encode(&idCategories)
}

func CreateCategories(w http.ResponseWriter, r *http.Request) {
	var newCategories models.Categories
	json.NewDecoder(r.Body).Decode(&newCategories)
	database.DB.Create(&newCategories)
	json.NewEncoder(w).Encode(&newCategories)
}*/
