package validation

import (
	"strings"
	"todo-list-api/app/model"
	"todo-list-api/exception"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateActivityGroup(request model.ActivityGroupRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Title, validation.Required.Error("title cannot be null")),
		validation.Field(&request.Email),
	)
	if err != nil {
		panic(exception.ValidationError{
			Message: strings.TrimSuffix(strings.Split(err.Error(), ": ")[1], "."),
		})
	}
}
