package validation

import (
	"todo-list-api/app/model"
	"todo-list-api/exception"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func ValidateActivityGroup(request model.ActivityGroupRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Email, validation.Required, is.Email),
		validation.Field(&request.Title, validation.Required),
	)
	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
