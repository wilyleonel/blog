package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wilyleo/apiPostgreSql/models"
)

func GetAllProfiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	profiles, err := models.GetAllProfiles()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(&profiles)
	}

}
func GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)["idpro"]
	idpro, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	profile, err := models.GetProfile(int(idpro))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(profile)

}
func CreateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var profile models.Profiles
	err := decoder.Decode(&profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = models.CreateProfile(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

/*func GetAllProfilesHandler(w http.ResponseWriter, r *http.Request) {
	var allProfile []models.Profile
	database.DB.Table("profiles").Find(&allProfile)

	json.NewEncoder(w).Encode(&allProfile)

}

func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	var idProfile models.Profile
	params := mux.Vars(r)
	database.DB.Find(&idProfile, params["id"])

	json.NewEncoder(w).Encode(&idProfile)
}

func CreateProfile(w http.ResponseWriter, r *http.Request) {
	var newProfile models.Profile
	json.NewDecoder(r.Body).Decode(&newProfile)
	database.DB.Create(&newProfile)
	json.NewEncoder(w).Encode(&newProfile)
}

func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	var profileDelete models.Profile
	params := mux.Vars(r)
	database.DB.First(&profileDelete, params["id"])
	if profileDelete.ID == 0 {
		w.WriteHeader(http.StatusNotFound) //error 400
		w.Write([]byte("user not found"))
	}

	database.DB.Delete(&profileDelete)
	w.WriteHeader(http.StatusOK)
}
*/
