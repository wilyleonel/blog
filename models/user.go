package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        int    `json:"id" gorm:"seriealizer:json;autoIncrement:true" json:"done"`
	Email     string //`json:"email"`
	Password  string `gorm:"<-"` //`json:"password"`
	ProfileID int    //`json:"perfil_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Profile   Profile `gorm:"foreignKey:ProfileID"`
}

//serializador

//
/*type User struct {
	Name [] byte                  `gorm:"serializer:json"`
	Roles Roles                   `gorm:"serializer:json"`
	Contracts    map [ string ] interface {} `gorm:"serializer:json"`
	JobInfo Job                     `gorm: "type:bytes;serializer:gob"`
	CreatedTime int64                   `gorm:"serializer:unixtime;type:time"`  // almacenar int como fecha y hora en la base de datos
  }

 escriba Roles [] cadena

 type Job struct {
   Título     cadena
	Ubicación cadena
	IsIntern bool
  }

 createdAt := time.Date( 2020 , 1 , 1 , 0 , 8 , 0 , 0 , time.UTC)
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

 DB.Create(&data)
*/
