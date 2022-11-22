package models

import (
	"time"

	"gorm.io/gorm"
)

type Postcateg struct {
	gorm.Model
	Id           int       `json:"id" gorm:"primarykey;autoIncrement:true" json:"done"`
	UpdatedAt    time.Time `gorm:"-:migration" gorm:"autoUpdateTime:false "`
	CreatedAt    time.Time `gorm:"-:migration" gorm:"autoCreateTime:false"`
	DeletedAt    time.Time //`gorm:"-:migration" gorm:"autoDeleteTime:false"`
	PostID       int
	CategoriesID int
	Post         Post       `gorm:"foreignKey:PostID"`
	Categories   Categories `gorm:"foreignKey:CategoriesID"`
}

/*createdAt := time.Date( 2020 , 1 , 1 , 0 , 8 , 0 , 0 , time.UTC)
data := User{
  Nombre: [] byte ( "jinzhu" ),
  Roles: [] string { "admin" , "propietario" },
  Contratos:    mapa [ cadena ] interfaz {}{ "nombre" : "jinzhu" , "edad" : 10 },
  CreatedTime: createdAt.Unix(),
  JobInfo: Job{
    Título:     "Desarrollador" ,
    Ubicación: "NY" ,
    IsIntern: falso ,
  },
}
DB.Create(&data) */
