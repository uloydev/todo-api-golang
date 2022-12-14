package initialize

import (
	"todo-list-api/app/controller"
)

var InitFunctions = []InitFunc{
	controller.InitializeActivityGroupController,
	controller.InitializeTodoController,
}
