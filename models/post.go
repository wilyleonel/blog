package models

import (
	"github.com/wilyleo/apiPostgreSql/database"
)

type Posts struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Title       string `json:"title"`
	IdUsers     int    `json:"idusers"`
	AllComments Comments
	Tipe        Cat
}
type NewPost struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Title       string `json:"title"`
	IdUsers     int    `json:"idusers"`
	Autor       NewProfile
}

func GetAllPosts() ([]NewPost, error) {
	query := `select p2.id,p2.description,p2.title,p2.idusers,p.name,p.lastname
	from posts as p2 inner join profiles as  p on p.id =idusers;`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var allPost []NewPost
	for rows.Next() {
		var post NewPost
		err := rows.Scan(
			&post.ID,
			&post.Description,
			&post.Title,
			&post.IdUsers,
			&post.Autor.Name,
			&post.Autor.Lastname,
		)
		if err != nil {
			return nil, err
		}
		allPost = append(allPost, post)
	}
	return allPost, nil
}

func GetPost(id int) ([]Posts, error) {
	query := `select p.id,p.description ,p.title,p.idusers,c.id,c.description,
    c.idusers,c.idposts,ca.category FROM posts p INNER JOIN "comments" as c
    ON c.idposts  = p.id LEFT JOIN categories ca ON ca.id = c.idposts where p.id =$1;`
	rows, err := database.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var postid []Posts
	for rows.Next() {
		var post Posts
		err := rows.Scan(
			&post.ID,
			&post.Description,
			&post.Title,
			&post.IdUsers,
			&post.AllComments.ID,
			&post.AllComments.Description,
			&post.AllComments.IdUsers,
			&post.AllComments.IdPosts,
			&post.Tipe.Category,
		)
		if err != nil {
			return nil, err
		}
		postid = append(postid, post)
	}
	return postid, nil
}

// func GetPost(id int) ([]Posts, error) {
// 	query := `select  p.id,p.description,p.title,p.idusers,c.id,c.description,c.idusers,c.idposts
// 	from posts  as p inner join comments as c on c.id =p.id  where p.id =$1;`
// 	rows, err := database.DB.Query(query, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	var postid []Posts
// 	for rows.Next() {
// 		var post Posts
// 		err := rows.Scan(
// 			&post.ID,
// 			&post.Description,
// 			&post.Title,
// 			&post.IdUsers,
// 			&post.AllComments.ID,
// 			&post.AllComments.Description,
// 			&post.AllComments.IdUsers,
// 			&post.AllComments.IdPosts,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}
// 		postid = append(postid, post)
// 	}
// 	return postid, nil
// }

/*func CreatePosts(post Posts) error {
	query := `insert into posts(description,title,idusers)values($1,$2,$3);`
	_, err := database.DB.Exec(query, post.Description, post.Title, post.IdUsers)
	if err != nil {
		return err
	}
	return nil
}*/
