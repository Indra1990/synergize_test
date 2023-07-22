package usecase

type TransactionService interface {
	CreateTransactionTopUp(cmd TransactionTopUpRequest) (err error)
}
