package models

import (
	"time"

	"gorm.io/gorm"
)

type Comments struct {
	gorm.Model
	Id          int    `json:"id" gorm:"primarykey;autoIncrement:true" json:"done"`
	Description string //`json:"descripcion"`
	UserID      int    //`json:"user_id"`
	PostID      int    //`json:"post_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        User `gorm:"foreignKey:UserID"`
	Post        Post `gorm:"foreignKey:PostID"`
}
