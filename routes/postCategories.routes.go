package routes

/*import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wilyleo/apiPostgreSql/database"
	"github.com/wilyleo/apiPostgreSql/models"
)

func GetAllpostCategories(w http.ResponseWriter, r *http.Request) {
	var allPostCategories []models.Postcateg
	database.DB.Table("postcategs").Find(&allPostCategories)
	json.NewEncoder(w).Encode(&allPostCategories)
}

func GetPostCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	var idPostCategories models.Postcateg
	params := mux.Vars(r)
	database.DB.First(&idPostCategories, params["id"])

	json.NewEncoder(w).Encode(&idPostCategories)
}*/
