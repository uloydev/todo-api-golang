package exception

import (
	"todo-list-api/app/model"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	_, ok := err.(ValidationError)
	if ok {
		return ctx.Status(400).JSON(model.WebResponse{
			Status:  "BAD_REQUEST",
			Message: err.Error(),
		})
	}

	return ctx.Status(500).JSON(model.WebResponse{
		Status:  "INTERNAL_SERVER_ERROR",
		Message: err.Error(),
	})
}
