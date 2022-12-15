package migrations

import (
	"todo-list-api/app/entity"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var CreateActivityGroupsTable *gormigrate.Migration = &gormigrate.Migration{
	ID: "create_activites_table",
	Migrate: func(dbConn *gorm.DB) error {
		return dbConn.AutoMigrate(&entity.Activity{})
	},
	Rollback: func(dbConn *gorm.DB) error {
		return dbConn.Migrator().DropTable("activities")
	},
}
