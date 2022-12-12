package models

import "github.com/wilyleo/apiPostgreSql/database"

type Profiles struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Img      string `json:"img"`
}
type NewProfile struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}



func GetAllProfiles() ([]Profiles, error) {
	var allProfiles []Profiles
	query := `select id,name,lastname,age,address,img from profiles;`
	rows, err := database.DB.Query(query)
	if err != nil {
		return allProfiles, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var lastname string
		var age int
		var address string
		var img string
		err := rows.Scan(&id, &name, &lastname, &age, &address, &img)
		if err != nil {
			return allProfiles, err
		}
		profile := Profiles{
			Id:       id,
			Name:     name,
			Lastname: lastname,
			Age:      age,
			Address:  address,
			Img:      img,
		}
		allProfiles = append(allProfiles, profile)
	}
	return allProfiles, nil
}
//idprofile
func GetProfile(idpro int) (Profiles, error) {
	var idprofile Profiles

	query := `select name,lastname,age,address,img from profiles where id=$1`
	row, err := database.DB.Query(query, idpro)
	if err != nil {
		return idprofile, err
	}

	defer row.Close()

	if row.Next() {
		var name string
		var lastname string
		var age int
		var address string
		var img string
		err := row.Scan(&name, &lastname, &age, &address, &img)
		if err != nil {
			return idprofile, err
		}

		idprofile = Profiles{
			Id:       idpro,
			Name:     name,
			Lastname: lastname,
			Age:      age,
			Address:  address,
			Img:      img,
		}
	}

	return idprofile, nil
}

func CreateProfile(profile Profiles) error {
	query := `insert into profiles(name,lastname,age,address,img) values($1, $2,$3,$4,$5);`
	_, err := database.DB.Exec(query, profile.Name, profile.Lastname, profile.Age, profile.Address, profile.Img)
	if err != nil {
		return err
	}
	return nil
}
