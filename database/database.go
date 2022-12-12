package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var DB *sql.DB

type connection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func Init() {
	err := godotenv.Load("config/.env")

	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err.Error())
		return
	}

	connInfo := connection{
		Host:     os.Getenv("POSTGRES_URL"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
	}

	// try to open our postgresql connection with our connection info
	DB, err = sql.Open("postgres", connToString(connInfo))
	if err != nil {
		fmt.Printf("Error connecting to the DB: %s\n", err.Error())
		return
	} else {
		fmt.Printf("DB is open\n")
	}

	// check if we can ping our DB
	err = DB.Ping()
	if err != nil {
		fmt.Printf("Error could not ping database: %s\n", err.Error())
		return
	} else {
		fmt.Printf("DB pinged successfully\n")
	}
}

// Take our connection struct and convert to a string for our db connection info
func connToString(info connection) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.User, info.Password, info.DBName)

}

// import (
// 	"database/sql"
// 	"log"

// 	_ "github.com/lib/pq"
// )

// var DB *sql.DB
// var DSN = "host=localhost user=postgres password=admin dbname=blog port=5432"

// func DBConnection() {
// 	var error error
// 	DB, error := sql.Open("postgres", DSN)
// 	if error != nil {
// 		log.Fatal(error)
// 	} else {
// 		log.Println("exito en la conexion a la base de datos ")
// 	}

// 	defer DB.Close()
// 	/*error = migration.MakeMigration(DB)
// 	if error != nil {
// 		log.Panic(error, "erejksdkkfjskf")

// 	}*/

// }
