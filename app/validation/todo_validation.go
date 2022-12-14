package validation

import (
	"todo-list-api/app/model"
	"todo-list-api/exception"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateTodo(request model.TodoRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.ActivityGroupId, validation.Required),
		validation.Field(&request.IsActive, validation.Required),
		validation.Field(&request.Title, validation.Required),
		validation.Field(&request.Priority, validation.Required),
	)
	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
