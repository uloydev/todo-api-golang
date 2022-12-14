package migrations

import (
	"todo-list-api/app/entity"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var CreateActivityGroupsTable *gormigrate.Migration = &gormigrate.Migration{
	ID: "create_activity_groups_table",
	Migrate: func(dbConn *gorm.DB) error {
		return dbConn.AutoMigrate(&entity.ActivityGroup{})
	},
	Rollback: func(dbConn *gorm.DB) error {
		return dbConn.Migrator().DropTable("activity_groups")
	},
}
