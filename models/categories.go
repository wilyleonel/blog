package models

import "gorm.io/gorm"

type Categories struct {
	gorm.Model
	Id   int    `json:"id" gorm:"primarykey;autoIncrement:true" json:"done"`
	Name string //`json:"nombre"`
}
