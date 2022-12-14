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

type ActivityGroupController struct {
	Service service.ActivityGroupService
}

func NewActivityGroupController(ActivityGroupService *service.ActivityGroupService) ActivityGroupController {
	return ActivityGroupController{Service: *ActivityGroupService}
}

func InitializeActivityGroupController(api *fiber.App, DB *gorm.DB) {
	ActivityGroupRepo := repository.NewActivityGroupRepository(DB)
	ActivityGroupService := service.NewActivityGroupService(ActivityGroupRepo.(*repository.ActivityGroupRepository))
	ActivityGroupController := NewActivityGroupController(ActivityGroupService.(*service.ActivityGroupService))
	ActivityGroupController.Route(api)
}

func (c *ActivityGroupController) Route(api *fiber.App) {
	api.Post("/activity-groups", c.Create)
	api.Get("/activity-groups", c.List)
	api.Get("/activity-groups/:id", c.FindById)
	api.Put("/activity-groups/:id", c.Update)
	api.Delete("/activity-groups/:id", c.Delete)
}

// CreateActivityGroup is a function to insert ActivityGroup to database
// @Summary      Create ActivityGroup
// @Description  Create New ActivityGroup / Register ActivityGroup
// @Tags         activity-groups
// @Accept       json
// @Produce      json
// @Param        ActivityGroup  body      model.ActivityGroupRequest  true  "Register ActivityGroup"
// @Success      2001  {object}  model.WebResponse{data=model.ActivityGroupResponse}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /activity-groups [post]
func (c *ActivityGroupController) Create(ctx *fiber.Ctx) error {
	var request model.ActivityGroupRequest

	err := ctx.BodyParser(&request)
	exception.PanicWhenError(err)

	resp := c.Service.Create(request)
	return ctx.Status(201).JSON(model.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    resp,
	})
}

// GetAllActivityGroup is a function to get all ActivityGroup data from database
// @Summary      Get All ActivityGroup
// @Description  get all ActivityGroup data from database
// @Tags         activity-groups
// @Accept       json
// @Produce      json
// @Success      200   {object}  model.WebResponse{data=[]model.ActivityGroupResponse}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /activity-groups [get]
func (c *ActivityGroupController) List(ctx *fiber.Ctx) error {
	var filter model.ActivityGroupFilter

	err := ctx.QueryParser(&filter)
	exception.PanicWhenError(err)

	resps := c.Service.List(&filter)

	return ctx.JSON(model.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    resps,
	})
}

// GeActivityGroup is a function to get ActivityGroup data by id from database
// @Summary      Get ActivityGroup by id
// @Description  get ActivityGroup data by id from database
// @Tags         activity-groups
// @Accept       json
// @Produce      json
// @Param id path int false "int valid" mininum(1)
// @Success      200   {object}  model.WebResponse{data=model.ActivityGroupResponse}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /activity-groups/ [get]
func (c *ActivityGroupController) FindById(ctx *fiber.Ctx) error {

	ID, err := strconv.Atoi(ctx.Params("id"))
	exception.PanicWhenError(err)

	resps := c.Service.FindById(uint(ID))

	return ctx.JSON(model.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    resps,
	})
}

// DeleteActivityGroup is a function to delete ActivityGroup to database
// @Summary      Delete ActivityGroup
// @Description  Delete ActivityGroup
// @Tags         activity-groups
// @Accept       json
// @Produce      json
// @Param id path int false "int valid" mininum(1)
// @Success      200   {object}  model.WebResponse{}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /activity-groups/{id} [delete]
func (controller *ActivityGroupController) Delete(c *fiber.Ctx) error {
	ID, err := strconv.Atoi(c.Params("id"))
	exception.PanicWhenError(err)

	controller.Service.DeleteById(uint(ID))
	return c.JSON(model.WebResponse{
		Status:  "Success",
		Message: "Success",
	})
}

// UpdateActivityGroup is a function to update ActivityGroup to database
// @Summary      Update ActivityGroup
// @Description  Update ActivityGroup
// @Tags         activity-groups
// @Accept       json
// @Produce      json
// @Param 		 id path int false "int valid" mininum(1)
// @Param        activity-groups  body      model.ActivityGroupRequest  true  "Update ActivityGroup"
// @Success      200   {object}  model.WebResponse{data=model.ActivityGroupResponse}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /activity-groups/{id} [put]
func (controller *ActivityGroupController) Update(c *fiber.Ctx) error {
	var request model.ActivityGroupRequest

	ID, err := strconv.Atoi(c.Params("id"))
	exception.PanicWhenError(err)

	err = c.BodyParser(&request)
	exception.PanicWhenError(err)

	resp := controller.Service.UpdateById(uint(ID), request)
	return c.JSON(model.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    resp,
	})
}
