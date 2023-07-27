package usecaseimpl

import (
	bankaccountrepository "synergize/api/bankaccount/repository"
	transactionrepository "synergize/api/transaction/repository"
	"synergize/entity"

	usecasetransaction "synergize/api/transaction/usecase"
)

type TransactionService struct {
	bankAccountRepository bankaccountrepository.BankAccountRepository
	transactionRepository transactionrepository.RepositoryTransaction
}

func NewTransactionService(
	transactionRepository transactionrepository.RepositoryTransaction,
	bankAccountRepository bankaccountrepository.BankAccountRepository,
) *TransactionService {
	return &TransactionService{
		transactionRepository: transactionRepository,
		bankAccountRepository: bankAccountRepository,
	}
}

func (trf *TransactionService) CreateTransactionTopUp(cmd usecasetransaction.TransactionTopUpRequest) (err error) {
	if err = validateTransactionTopUp(cmd); err != nil {
		return
	}

	bankAccount, bankAccountErr := trf.bankAccountRepository.FindBankAccountByUserId(cmd.UserId)
	if bankAccountErr != nil {
		err = bankAccountErr
		return
	}

	transaction := entity.Transaction{
		Type:          "TOPUP",
		Status:        "SUCCESS",
		Notes:         cmd.Notes,
		Amount:        cmd.Amount,
		UserId:        cmd.UserId,
		AccountBankId: bankAccount.ID,
	}

	if err = trf.transactionRepository.CreateTransaction(&transaction); err != nil {
		return
	}

	return
}
