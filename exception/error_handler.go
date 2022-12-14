package exception

import (
	"todo-list-api/app/model"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var code int
	resp := model.WebResponse{Message: err.Error()}

	if _, ok := err.(ValidationError); ok {
		code = 400
		resp.Status = "Bad Request"
	} else if _, ok := err.(NotFoundError); ok {
		code = 404
		resp.Status = "Not Found"
	} else {
		code = 500
		resp.Status = "Internal Server Error"
	}

	return ctx.Status(code).JSON(resp)
}
