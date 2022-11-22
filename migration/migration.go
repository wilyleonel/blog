package migration

import (
	"github.com/wilyleo/apiPostgreSql/database"
	"github.com/wilyleo/apiPostgreSql/models"
)

func Migration() {
	database.DB.Migrator().DropTable(models.Post{})
	database.DB.Migrator().DropTable(models.Comments{})
	database.DB.Migrator().DropTable(models.Categories{})
	database.DB.Migrator().DropTable(models.Postcateg{})
	database.DB.Migrator().DropTable(models.Profile{})
	database.DB.Migrator().DropTable(models.User{})

	database.DB.AutoMigrate(models.Post{})
	database.DB.AutoMigrate(models.Comments{})
	database.DB.AutoMigrate(models.Categories{})
	database.DB.AutoMigrate(models.Postcateg{})
	database.DB.AutoMigrate(models.Profile{})
	database.DB.AutoMigrate(models.User{})
}
