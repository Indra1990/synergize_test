package usecaseimpl

import (
	transactionrepository "synergize/api/transaction/repository"
	usecasetransaction "synergize/api/transaction/usecase"
	"synergize/entity"
)

type TransactionService struct {
	transactionRepository transactionrepository.RepositoryTransaction
}

func NewTransactionService(transactionRepository transactionrepository.RepositoryTransaction) *TransactionService {
	return &TransactionService{
		transactionRepository: transactionRepository,
	}
}

func (trf *TransactionService) CreateTransactionTopUp(cmd usecasetransaction.TransactionTopUpRequest) (err error) {
	if err = validateTransactionTopUp(cmd); err != nil {
		return
	}
	transaction := entity.Transaction{
		Type:          "TOPUP",
		Status:        "SUCCESS",
		Notes:         cmd.Notes,
		Amount:        cmd.Amount,
		UserId:        cmd.UserId,
		AccountBankId: cmd.BankAccountId,
	}

	if err = trf.transactionRepository.CreateTransaction(&transaction); err != nil {
		return
	}

	return
}
