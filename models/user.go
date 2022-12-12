package models

import (
	"fmt"


	"github.com/wilyleo/apiPostgreSql/bcrypt"
	"github.com/wilyleo/apiPostgreSql/database"
)

type Users struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IdProfile int    `json:"idprofiles"`
	Profile   Profiles
}
type NewUser struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IdProfile int    `json:"idprofiles"`
	User      NewProfile
}

func GetAllUsers() ([]Users, error) {
	query := `select p.id,p.name,p.lastname,p.age,p.address,p.img,
	u.id,u.email,u.password,u.idprofiles from users as u inner join profiles as p on p.id=u.idprofiles;`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var userss []Users
	for rows.Next() {
		var user Users
		err := rows.Scan(
			&user.Profile.Id,
			&user.Profile.Name,
			&user.Profile.Lastname,
			&user.Profile.Age,
			&user.Profile.Address,
			&user.Profile.Img,
			&user.ID,
			&user.Email,
			&user.Password,
			&user.IdProfile)
		if err != nil {
			return nil, err
		}
		fmt.Sprintln(user, "-.sd-as.d-as")
		userss = append(userss, user)
	}
	return userss, nil
}

func GetUsers(id int) ([]Users, error) {
	query := `select u.id,u.email,u.password,u.idprofiles,p.id ,p.name,p.lastname,p.age ,p.address ,p.img
	from users as u inner join profiles as p on p.id=idprofiles  where u.id  =$1
	;`
	rows, err := database.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var iduser []Users
	for rows.Next() {
		var user Users
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			&user.IdProfile,
			&user.Profile.Id,
			&user.Profile.Name,
			&user.Profile.Lastname,
			&user.Profile.Age,
			&user.Profile.Address,
			&user.Profile.Img,
		)
		if err != nil {
			return nil, err
		}
		iduser = append(iduser, user)
	}
	return iduser, nil
}

/*func GetAllUsers() ([]Users, error) {
	var allUsers []Users
	query := `select * from users inner join profiles on users.id = profiles.idpro;`
	rows, err := database.DB.Query(query)
	fmt.Println(rows)
	fmt.Println(query, "query")
	if err != nil {
		fmt.Println("error de prueba")
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var email string
		var password string
		var idprofiles int
		var idpro int
		var name string
		var lastname string
		var age int
		var address string
		var img string
		err := rows.Scan(&id, &email, &password, &idprofiles, &idpro, &name, &lastname, &age, &address, &img)
		if err != nil {
			return allUsers, err
		}
		users1 := Users{
			ID:        id,
			Email:     email,
			Password:  password,
			IdProfile: idprofiles,
		}
		profile := Profiles{
			IdPro:    idpro,
			Name:     name,
			Lastname: lastname,
			Age:      age,
			Address:  address,
			Img:      img,
		}
		allUsers = append(allUsers, users1)
		fmt.Println(users1)
		fmt.Println(profile, "--------------")

	}
	return allUsers, nil
}*/

// func GetUser(id int) (Users, error) {
// 	var userId Users

// 	query := `select email,password,idprofiles from users where id=$1`
// 	row, err := database.DB.Query(query, id)
// 	if err != nil {
// 		return userId, err
// 	}

// 	defer row.Close()

// 	if row.Next() {
// 		var idprofiles int
// 		var email, password string

// 		err := row.Scan(&email, &password, &idprofiles)
// 		if err != nil {
// 			return userId, err
// 		}

// 		userId = Users{
// 			ID:        id,
// 			Email:     email,
// 			Password:  password,
// 			IdProfile: idprofiles,
// 		}

//		}
//		return userId, nil
//	}
func CreateUser(user Users, err error) error {
	user.Password, err = bcrypt.HashPassword(user.Password)
	if err != nil {
		return err
	}

	query := `insert into users(email,password,idprofiles) values($1, $2,$3);`

	_, err = database.DB.Exec(query, user.Email, user.Password, user.IdProfile)
	if err != nil {
		return err
	}
	return nil
}
