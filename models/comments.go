package models

import (
	"github.com/wilyleo/apiPostgreSql/database"
)

type Comments struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	IdUsers     int    `json:"idusers"`
	IdPosts     int    `json:"idposts"`
}

func GetAllComments() ([]Comments, error) {
	var allComments []Comments
	query := `select id,description,idusers,idposts from comments;`
	rows, err := database.DB.Query(query)
	if err != nil {
		return allComments, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var description string
		var idusers int
		var idposts int
		err := rows.Scan(&id, &description, &idusers, &idposts)
		if err != nil {
			return allComments, err
		}
		comment := Comments{
			ID:          id,
			Description: description,
			IdUsers:     idusers,
			IdPosts:     idposts,
		}
		allComments = append(allComments, comment)
	}
	return allComments, nil
}

func GetComments(id int) (Comments, error) {
	var idComments Comments
	query := `select description,idusers,idposts from comments;`
	row, err := database.DB.Query(query, id)
	if err != nil {
		return idComments, err
	}
	defer row.Close()
	if row.Next() {
		var description string
		var idusers int
		var idposts int
		err := row.Scan(&description)
		if err != nil {
			return idComments, err
		}
		idComments = Comments{
			ID:          id,
			Description: description,
			IdUsers:     idusers,
			IdPosts:     idposts,
		}
	}
	return idComments, nil
}
func CreateComments(comment Comments) error {
	query := `insert into comments(description,idusers,idposts)values($1,$2,$3);`
	_, err := database.DB.Exec(query, comment.Description, comment.IdUsers, comment.IdPosts)
	if err != nil {
		return err
	}
	return nil
}
