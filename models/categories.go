package models

import (
	"github.com/wilyleo/apiPostgreSql/database"
)

type Categories struct {
	ID       int    `json:"id"`
	Category string `json:"category"`
}
type Cat struct {
	Category string `json:"category"`
}

func GetAllCategories() ([]Categories, error) {
	var allCategories []Categories
	query := `select id,category from categories;`
	rows, err := database.DB.Query(query)
	if err != nil {
		return allCategories, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var category string
		err := rows.Scan(&id, &category)
		if err != nil {
			return allCategories, err
		}
		categorie := Categories{
			ID:       id,
			Category: category,
		}
		allCategories = append(allCategories, categorie)
	}
	return allCategories, nil
}
func GetCategory(id int) (Categories, error) {
	var idcategories Categories
	query := `select category from categories where id=$1`
	row, err := database.DB.Query(query, id)
	if err != nil {
		return idcategories, err
	}
	defer row.Close()
	if row.Next() {
		var category string
		err := row.Scan(&category)
		if err != nil {
			return idcategories, err
		}
		idcategories = Categories{
			ID:       id,
			Category: category,
		}
	}
	return idcategories, nil
}
func CreateCategories(category Categories) error {
	query := `insert into categories(name)values($1);`
	_, err := database.DB.Exec(query, category)
	if err != nil {
		return err
	}
	return nil
}
