package controller

import (
	"strconv"
	"todo-list-api/app/model"
	"todo-list-api/app/repository"
	"todo-list-api/app/service"
	"todo-list-api/exception"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TodoController struct {
	Service service.TodoService
}

func NewTodoController(TodoService *service.TodoService) TodoController {
	return TodoController{Service: *TodoService}
}

func InitializeTodoController(api *fiber.App, DB *gorm.DB) {
	TodoRepo := repository.NewTodoRepository(DB)
	TodoService := service.NewTodoService(TodoRepo.(*repository.TodoRepository))
	TodoController := NewTodoController(TodoService.(*service.TodoService))
	TodoController.Route(api)
}

func (c *TodoController) Route(api *fiber.App) {
	api.Post("/todo-items", c.Create)
	api.Get("/todo-items", c.List)
	api.Get("/todo-items/:id", c.FindById)
	api.Put("/todo-items/:id", c.Update)
	api.Delete("/todo-items/:id", c.Delete)
}

// CreateTodo is a function to insert Todo to database
// @Summary      Create Todo
// @Description  Create New Todo / Register Todo
// @Tags         todo-items
// @Accept       json
// @Produce      json
// @Param        Todo  body      model.TodoRequest  true  "Register Todo"
// @Success      201   {object}  model.WebResponse{data=model.TodoResponse}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /todo-items [post]
func (c *TodoController) Create(ctx *fiber.Ctx) error {
	var req model.TodoRequest

	err := ctx.BodyParser(&req)
	exception.PanicWhenError(err)

	resp := c.Service.Create(req)
	return ctx.Status(201).JSON(model.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    resp,
	})
}

// GetAllTodo is a function to get all Todo data from database
// @Summary      Get All Todo
// @Description  get all Todo data from database
// @Tags         todo-items
// @Accept       json
// @Produce      json
// @Param activity_group_id query int false "int valid" mininum(1)
// @Success      200   {object}  model.WebResponse{data=[]model.TodoResponse}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /todo-items [get]
func (c *TodoController) List(ctx *fiber.Ctx) error {
	var filter model.TodoFilter

	err := ctx.QueryParser(&filter)
	exception.PanicWhenError(err)

	resps := c.Service.List(&filter)

	return ctx.JSON(model.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    resps,
	})
}

// GeTodo is a function to get Todo data by id from database
// @Summary      Get Todo by id
// @Description  get Todo data by id from database
// @Tags         todo-items
// @Accept       json
// @Produce      json
// @Param id path int false "int valid" mininum(1)
// @Success      200   {object}  model.WebResponse{data=model.TodoResponse}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /todo-items/ [get]
func (c *TodoController) FindById(ctx *fiber.Ctx) error {

	ID, err := strconv.Atoi(ctx.Params("id"))
	exception.PanicWhenError(err)

	resps := c.Service.FindById(uint(ID))

	return ctx.JSON(model.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    resps,
	})
}

// DeleteTodo is a function to delete Todo to database
// @Summary      Delete Todo
// @Description  Delete Todo
// @Tags         todo-items
// @Accept       json
// @Produce      json
// @Param id path int false "int valid" mininum(1)
// @Success      200   {object}  model.WebResponse{}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /todo-items/{id} [delete]
func (controller *TodoController) Delete(c *fiber.Ctx) error {
	ID, err := strconv.Atoi(c.Params("id"))
	exception.PanicWhenError(err)

	controller.Service.DeleteById(uint(ID))
	return c.JSON(model.WebResponse{
		Status:  "Success",
		Message: "Success",
	})
}

// UpdateTodo is a function to update Todo to database
// @Summary      Update Todo
// @Description  Update Todo
// @Tags         todo-items
// @Accept       json
// @Produce      json
// @Param 		 id path int false "int valid" mininum(1)
// @Param        todo-items  body      model.TodoRequest  true  "Update Todo"
// @Success      200   {object}  model.WebResponse{data=model.TodoResponse}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /todo-items/{id} [put]
func (controller *TodoController) Update(c *fiber.Ctx) error {
	var req model.TodoRequest

	ID, err := strconv.Atoi(c.Params("id"))
	exception.PanicWhenError(err)

	err = c.BodyParser(&req)
	exception.PanicWhenError(err)

	resp := controller.Service.UpdateById(uint(ID), req)
	return c.JSON(model.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    resp,
	})
}
