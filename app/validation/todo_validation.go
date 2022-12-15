package validation

import (
	"strings"
	"todo-list-api/app/model"
	"todo-list-api/exception"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateCreateTodo(request model.TodoRequest) {
	validation.ErrRequired = validation.ErrRequired.SetMessage("")
	err := validation.ValidateStruct(&request,
		validation.Field(&request.ActivityGroupId, validation.Required.Error("activity_group_id cannot be null")),
		validation.Field(&request.Title, validation.Required.Error("title cannot be null")),
	)
	if err != nil {
		panic(exception.ValidationError{
			Message: strings.TrimSuffix(strings.Split(err.Error(), ": ")[1], "."),
		})
	}
}

func ValidateUpdateTodo(request model.TodoRequest) {
	validation.ErrRequired = validation.ErrRequired.SetMessage("")
	err := validation.ValidateStruct(&request,
		validation.Field(&request.ActivityGroupId),
		validation.Field(&request.IsActive),
	)
	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
