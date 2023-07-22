package usecaseimpl

import (
	"synergize/api/transaction/usecase"

	validation "github.com/go-ozzo/ozzo-validation"
)

func validateTransactionTopUp(cmd usecase.TransactionTopUpRequest) (err error) {
	err = validation.ValidateStruct(&cmd,
		validation.Field(&cmd.Amount, validation.Required),
	)

	if err != nil {
		return
	}

	return
}
