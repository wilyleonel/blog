package main

import (
	"github.com/wilyleo/apiPostgreSql/database"
	"github.com/wilyleo/apiPostgreSql/handlers"
	"github.com/wilyleo/apiPostgreSql/migration"
	_ "github.com/wilyleo/apiPostgreSql/migration"
	"github.com/wilyleo/apiPostgreSql/seeds"
	_ "github.com/wilyleo/apiPostgreSql/seeds"
)

func main() {

	database.Init()
	migration.MakeMigration()
	seeds.RegisterProfiles()
	seeds.RegisterUser()
	seeds.RegisterPosts()
	seeds.RegisterComments()
	seeds.RegisterCategories()
	/*seeds.NewPostsCategories()*/
	handlers.Rutas()

}
