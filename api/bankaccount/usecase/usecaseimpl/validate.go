package usecaseimpl

import (
	"errors"
	"fmt"
	"synergize/api/bankaccount/usecase"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

const (
	BCA     string = "BCA"
	BNI     string = "BNI"
	BRI     string = "BRI"
	MANDIRI string = "MANDIRI"
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

	if !validateBankAllowed(cmd.BankName) {
		anyBankName := fmt.Sprintf("Bank name just allow %s,%s,%s,%s", BCA, BNI, BRI, MANDIRI)
		err = errors.New(anyBankName)
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

	if !validateBankAllowed(cmd.BankName) {
		anyBankName := fmt.Sprintf("Bank name just allow %s,%s,%s,%s", BCA, BNI, BRI, MANDIRI)
		err = errors.New(anyBankName)
		return
	}

	return
}

func validateBankAllowed(bankName string) bool {
	var checkBankName bool
	switch bankName {
	case BCA:
		checkBankName = true
	case BNI:
		checkBankName = true
	case BRI:
		checkBankName = true
	case MANDIRI:
		checkBankName = true
	default:
		checkBankName = false
	}

	return checkBankName
}
