package routs

import (
	"net/http"
)

func Homehandler(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("hello wAsadszsd1213sSord,hola mundo")))
}
