package usecaseimpl

import (
	"strings"
	"synergize/api/bankaccount/repository"
	"synergize/api/bankaccount/usecase"
	"synergize/entity"
)

type BankAccountService struct {
	repo repository.BankAccountRepository
}

func NewBankAccountService(repo repository.BankAccountRepository) *BankAccountService {
	return &BankAccountService{repo: repo}
}

func (b *BankAccountService) Create(cmd usecase.BankAccountRequest) (err error) {
	if err = validateBankAccount(cmd); err != nil {
		return
	}

	if err = b.repo.CheckUserIdExist(cmd.UserId); err != nil {
		return
	}

	ent := entity.BankAccount{
		UserId:            cmd.UserId,
		AccountName:       cmd.AccountName,
		BankName:          strings.ToUpper(cmd.BankName),
		AccountBankNumber: cmd.AccountBankNumber,
	}

	if err = b.repo.Create(&ent); err != nil {
		return
	}

	return
}

func (b *BankAccountService) FindbyId(bankAccountId int) (dto usecase.BankAccountResponse, err error) {
	bankAccount, bankAccountErr := b.repo.FindbyId(bankAccountId)
	if bankAccountErr != nil {
		err = bankAccountErr
		return
	}

	dto = usecase.BankAccountResponse{
		UserId:            bankAccount.UserId,
		BankName:          bankAccount.BankName,
		AccountName:       bankAccount.AccountName,
		AccountBankNumber: bankAccount.AccountBankNumber,
		BankAccountUserResponse: &usecase.BankAccountUserResponse{
			Username:    bankAccount.User.Username,
			Email:       bankAccount.User.Email,
			PhoneNumber: bankAccount.User.PhoneNumber,
		},
	}
	return
}

func (b *BankAccountService) Update(cmd usecase.BankAccountRequestUpdate) (err error) {
	if err = validateBankAccountUpdate(cmd); err != nil {
		return
	}

	bankAccount, bankAccountErr := b.repo.FindbyId(int(cmd.ID))
	if bankAccountErr != nil {
		err = bankAccountErr
		return
	}

	ent := entity.BankAccount{
		ID:                bankAccount.ID,
		UserId:            cmd.UserId,
		BankName:          strings.ToUpper(cmd.BankName),
		AccountName:       bankAccount.AccountName,
		AccountBankNumber: cmd.AccountBankNumber,
	}

	if err = b.repo.Update(&ent); err != nil {
		return
	}

	return
}
