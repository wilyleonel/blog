package seeds

import (
	"github.com/wilyleo/apiPostgreSql/database"
)

// }
func RegisterProfiles() {
	profile1 := `INSERT INTO profiles  VALUES ('1','wilyuno', 'Juan', 342, 'avenida 01', 'qwedsas1');`
	profile2 := `INSERT INTO profiles  VALUES ('2','wilydos', 'Juan', 342, 'avenida 012', 'qwedsas2');`
	profile3 := `INSERT INTO profiles  VALUES ('3','wilytres', 'Juan', 342, 'avenida 013', 'qwedsas3');`
	profile4 := `INSERT INTO profiles  VALUES ('4','wilycuatro', 'Juan', 342, 'avenida 014', 'qwedsas4');`
	profile5 := `INSERT INTO profiles  VALUES ('5','wilycinco', 'Juan', 342, 'avenida 015', 'qwedsas5');`
	database.DB.Exec(profile1)
	database.DB.Exec(profile2)
	database.DB.Exec(profile3)
	database.DB.Exec(profile4)
	database.DB.Exec(profile5)
}
func RegisterUser() {
	user1 := `INSERT INTO users  VALUES ('1','email1', 'password1', '1');`
	user2 := `INSERT INTO users  VALUES ('2','email2', 'password2', '1');`
	user3 := `INSERT INTO users  VALUES ('3','email3', 'password3', '1');`
	user4 := `INSERT INTO users  VALUES ('4','email4', 'password4', '4');`
	user5 := `INSERT INTO users  VALUES ('5','email5', 'password5', '5');`
	database.DB.Exec(user1)
	database.DB.Exec(user2)
	database.DB.Exec(user3)
	database.DB.Exec(user4)
	database.DB.Exec(user5)
}
func RegisterPosts() {
	post1 := `INSERT INTO posts  VALUES ('1','description posts1', 'title posts1', '1');`
	post2 := `INSERT INTO posts  VALUES ('2','description posts2', 'title posts2', '1');`
	post3 := `INSERT INTO posts  VALUES ('3','description posts3', 'title posts3', '1');`
	post4 := `INSERT INTO posts  VALUES ('4','description posts4', 'title posts4', '1');`
	post5 := `INSERT INTO posts  VALUES ('5','description posts5', 'title posts5', '5');`
	database.DB.Exec(post1)
	database.DB.Exec(post2)
	database.DB.Exec(post3)
	database.DB.Exec(post4)
	database.DB.Exec(post5)
}
func RegisterComments() {
	comment1 := `INSERT INTO comments  VALUES ('1','description comments1', '1', '1');`
	comment2 := `INSERT INTO comments  VALUES ('2','description comments2', '2', '2');`
	comment3 := `INSERT INTO comments  VALUES ('3','description comments3', '2', '3');`
	comment4 := `INSERT INTO comments  VALUES ('4','description comments4', '2', '4');`
	comment5 := `INSERT INTO comments  VALUES ('5','description comments5', '1', '5');`
	database.DB.Exec(comment1)
	database.DB.Exec(comment2)
	database.DB.Exec(comment3)
	database.DB.Exec(comment4)
	database.DB.Exec(comment5)
}
func RegisterCategories() {
	cate1 := `INSERT INTO categories  VALUES ('1','nameCategories1');`
	cate2 := `INSERT INTO categories  VALUES ('2','nameCategories2');`
	cate3 := `INSERT INTO categories  VALUES ('3','nameCategories3');`
	cate4 := `INSERT INTO categories  VALUES ('4','nameCategories4');`
	cate5 := `INSERT INTO categories  VALUES ('5','nameCategories5');`
	database.DB.Exec(cate1)
	database.DB.Exec(cate2)
	database.DB.Exec(cate3)
	database.DB.Exec(cate4)
	database.DB.Exec(cate5)
}

// import (
// 	"github.com/wilyleo/apiPostgreSql/database"
// 	"github.com/wilyleo/apiPostgreSql/models"
// )

// func RegisterUser() {
// 	var usuario = []models.Users{{Email: "dasdkaj@gmail.com", Password: "admaisdonisad"},
// 		{Email: "dasdkaj@gmail.com", Password: "admaisdonisad"},
// 		{Email: "dasdkaj@gmail.com", Password: "admaisdonisad"}}
// 	database.DB.Exec(&usuario)

/*func NewPosts() {
	var posts = []models.Post{{Description: "esto es una descripcion de la seccion de publicaciones", Title: "new posts", UserID: 1, CommentsID: 1},
		{Description: "esto es una descripadsdcion de la seccion de publicaciones", Title: "new podsadsts", UserID: 2, CommentsID: 2},
		{Description: "esto es una descripadsdcion de la seccion de publicaciones", Title: "new podsadsts", UserID: 1, CommentsID: 2},
		{Description: "esto es una descripadsdcion de la seccion de publicaciones", Title: "new podsadsts", UserID: 1, CommentsID: 1}}
	database.DB.Create(&posts)
}
func NewComments() {
	var comments = []models.Comments{{Description: "esto es una descripcion de la seccion comentarios", UserID: 1, PostID: 1},
		{Description: "esto es una descripcion de la seccion comentarios", UserID: 1, PostID: 2}}
	database.DB.Create(&comments)
}*/

/*func NewCategories() {
	var categories = []models.Categories{{Name: "new name the categories", PostID: 2}, {Name: "new name the categories", PostID: 1}}
	database.DB.Create(&categories)
}
func NewPostsCategories() {
	var postsCategories = []models.Postcateg{{PostID: 1, CategoriesID: 2},
		{PostID: 2, CategoriesID: 1}}
	database.DB.Create(&postsCategories)
}*/
