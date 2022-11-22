package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Id          int    `json:"id" gorm:"primarykey;autoIncrement:true" json:"done"`
	Description string //`json:"descripcion"`
	Title       string //`json:"title"`
	UserID      int    //`json:"user_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        User `gorm:"foreignKey:UserID"`
}
