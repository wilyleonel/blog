package routes

import (
	"encoding/json"
	"net/http"
)

/*
	func Homehandler(w http.ResponseWriter, r *http.Request) {
		//w.Write(([]byte("hello wAsadszsd1213sSord,hola mundo")))
	}
*/
type HomeResponse struct {
	Message string `json:"message"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HomeResponse{
		Message: "welcome to blog",
	})

}
