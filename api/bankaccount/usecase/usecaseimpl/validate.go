package usecaseimpl

import (
	"synergize/api/bankaccount/usecase"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func validateBankAccount(cmd usecase.BankAccountRequest) (err error) {
	err = validation.ValidateStruct(&cmd,
		validation.Field(&cmd.BankName, validation.Required),
		validation.Field(&cmd.AccountName, validation.Required),
		validation.Field(&cmd.AccountBankNumber, validation.Required, is.Digit, validation.Length(8, 20)),
	)

	if err != nil {
		return
	}

	return
}

func validateBankAccountUpdate(cmd usecase.BankAccountRequestUpdate) (err error) {
	err = validation.ValidateStruct(&cmd,
		validation.Field(&cmd.BankName, validation.Required),
		validation.Field(&cmd.AccountName, validation.Required),
		validation.Field(&cmd.AccountBankNumber, validation.Required, is.Digit, validation.Length(8, 20)),
	)

	if err != nil {
		return
	}

	return
}
