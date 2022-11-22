package models

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	Id        int    `json:"id" gorm:"primarykey;autoIncrement:true" json:"done"`
	Name      string `faker:"name"`
	lastname  string //`json:"firsname"`
	Age       int    //`json:"edad"`
	Address   string //`json:"direccion"`
	Img       string //`json:"img"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
