package usecase

type BankAccountService interface {
	Create(cmd BankAccountRequest) (err error)
	FindbyId(bankAccountId int) (dto BankAccountResponse, err error)
	Update(cmd BankAccountRequestUpdate) (err error)
}
