package seeds

import (
	"github.com/bxcodec/faker/v3"
)

func (s Seed) UserSeed() {

	for i := 0; i < 10; i++ {
		//prepare the statement
		stmt, _ := s.database.Prepare(`INSERT INTO user(name, email,password) VALUES (?,?,?)`)
		// execute query
		_, err := stmt.Exec(faker.Name(), faker.Email(), faker.Password())
		if err != nil {
			panic(err)
		}
	}
}
