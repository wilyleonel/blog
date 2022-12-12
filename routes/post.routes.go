package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wilyleo/apiPostgreSql/models"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post, err := models.GetAllPosts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(&post)
	}
}

// posts/id
func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	post, err := models.GetPost(int(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(post)
}

//  //create pots
// func CreatePost(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	decoder := json.NewDecoder(r.Body)
// 	var post models.Posts
// 	err := decoder.Decode(&post)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
// 	err = models.CreatePosts(post)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(err.Error()))
// 		return
// 	} else {
// 		w.WriteHeader(http.StatusOK)
// 	}
// }

// // user/id
// /*func GetAllPostsHandler(w http.ResponseWriter, r *http.Request) {
// 	var allPost []models.Post
// 	database.DB.Table("posts").Joins("Comments").Find(&allPost)
// 	json.NewEncoder(w).Encode(&allPost)
// }

// func CreatePost(w http.ResponseWriter, r *http.Request) {
// 	var newPost models.Post
// 	json.NewDecoder(r.Body).Decode(&newPost)

// 	createPost := database.DB.Create(&newPost)
// 	err := createPost.Error
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest) //error 400
// 		w.Write([]byte(err.Error()))
// 	}
// 	json.NewEncoder(w).Encode(&newPost)
// }
// func GetPostHandler(w http.ResponseWriter, r *http.Request) {
// 	var idPost models.Post
// 	params := mux.Vars(r)
// 	database.DB.First(&idPost, params["id"])

// 	json.NewEncoder(w).Encode(&idPost)
// }
// func PostDelete(w http.ResponseWriter, r *http.Request) {
// 	var postDelete models.Post
// 	params := mux.Vars(r)
// 	database.DB.First(&postDelete, params["id"])
// 	if postDelete.ID == 0 {
// 		w.WriteHeader(http.StatusNotFound) //error 400
// 		w.Write([]byte("user not found"))
// 	}

// 	database.DB.Delete(&postDelete)
// 	w.WriteHeader(http.StatusOK)
// }
// */
