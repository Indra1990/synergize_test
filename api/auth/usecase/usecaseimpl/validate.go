package usecaseimpl

import (
	"synergize/api/auth/usecase"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func validateRegister(cmd usecase.RegisterRequest) (err error) {
	err = validation.ValidateStruct(&cmd,
		validation.Field(&cmd.Username, validation.Required, validation.Length(3, 0)),
		validation.Field(&cmd.Email, validation.Required, is.Email),
		validation.Field(&cmd.Password, validation.Required, validation.Length(6, 0)),
		validation.Field(&cmd.PhoneNumber, validation.Required, is.Digit),
	)

	if err != nil {
		return
	}

	return
}
