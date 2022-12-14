package exception

import (
	"todo-list-api/app/model"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	_, ok := err.(ValidationError)
	if ok {
		return ctx.JSON(model.WebResponse{
			Status:  "BAD_REQUEST",
			Message: err.Error(),
			Data:    map[string]any{},
		})
	}

	return ctx.JSON(model.WebResponse{
		Status:  "INTERNAL_SERVER_ERROR",
		Message: err.Error(),
		Data:    map[string]any{},
	})
}
