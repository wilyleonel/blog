package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wilyleo/apiPostgreSql/routes"
)

func Rutas() {
	r := mux.NewRouter()
	// 	// 	//login users
	// 	// 	//r.HandleFunc("/signup", routes.SignUp).Methods("POST")
	// 	// 	//r.HandleFunc("/signin", routes.SignIn).Methods("POST")

	r.HandleFunc("/", routes.Home)

	//ruotes.use
	r.HandleFunc("/users", routes.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUser).Methods("GET")
	r.HandleFunc("/users", routes.CreateUser).Methods("POST")
	// 	// r.HandleFunc("/users/{id}", routes.DeleteUser).Methods("DELETE")

	// 	// //routes.Posts
	r.HandleFunc("/posts", routes.GetAllPosts).Methods("GET")
	r.HandleFunc("/posts/{id}", routes.GetPost).Methods("GET")
	// r.HandleFunc("/users/{id}/posts/{post}", routes.CreatePost).Methods("GET")
	// 	// r.HandleFunc("/posts/{id}", routes.PostDelete).Methods("DELETE")

	// 	// //routes.Comments
	r.HandleFunc("posts/{id}/comments", routes.GetAllComments).Methods("GET")
	r.HandleFunc("/users/comments/{id}", routes.GetComment).Methods("GET")
	r.HandleFunc("user/comments", routes.CreateComments).Methods("POST")
	// 	// r.HandleFunc("/users/{1}/comments", routes.DeleteUser).Methods("DELETE")

	// 	// //routes.profile
	r.HandleFunc("/profiles", routes.GetAllProfiles).Methods("GET")
	r.HandleFunc("/users/{id}/profiles", routes.GetProfile).Methods("GET")
	r.HandleFunc("/profiles", routes.CreateProfile).Methods("POST")
	// 	// r.HandleFunc("/users/{id}/profiles/{profile}", routes.DeleteProfile).Methods("DELETE")

	// 	// //routes.categories
	r.HandleFunc("/categories", routes.GetAllCategories).Methods("GET")
	r.HandleFunc("/categories", routes.GetCategories).Methods("GET")
	r.HandleFunc("/categories", routes.CreateCategories).Methods("GET")

	// 	// //routes..post/comments
	// 	// r.HandleFunc("/postcategories", routs.GetAllpostCategories).Methods("GET")
	// 	// r.HandleFunc("/postcategories/{id}", routs.GetPostCategoriesHandler).Methods("GET")

	// 	// //routes.categories
	// 	// //r.HandleFunc("user/")
	http.ListenAndServe(":3000", r)
}
