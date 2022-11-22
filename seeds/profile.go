package seeds

import (
	"github.com/bxcodec/faker/v3"
)

func (s Seed) ProfileSeed() {

	for i := 0; i < 10; i++ {
		//prepare the statement
		stmt, _ := s.database.Prepare(`INSERT INTO user(name,firstname,address,img) VALUES (?,?,?,?,?)`)
		// execute query
		_, err := stmt.Exec(faker.Name(), faker.LastName(), faker.GetAddress().Latitude())
		if err != nil {
			panic(err)
		}
	}
}
