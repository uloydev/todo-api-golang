package db

import (
	"todo-list-api/db/migrations"

	"github.com/go-gormigrate/gormigrate/v2"
)

var Migrations = []*gormigrate.Migration{
	migrations.CreateActivityGroupsTable,
	migrations.CreateTodosTable,
}
