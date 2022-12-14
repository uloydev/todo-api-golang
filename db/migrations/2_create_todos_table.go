package migrations

import (
	"todo-list-api/app/entity"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var CreateTodosTable *gormigrate.Migration = &gormigrate.Migration{
	ID: "create_todos_table",
	Migrate: func(dbConn *gorm.DB) error {
		return dbConn.AutoMigrate(&entity.Todo{})
	},
	Rollback: func(dbConn *gorm.DB) error {
		return dbConn.Migrator().DropTable("todos")
	},
}
