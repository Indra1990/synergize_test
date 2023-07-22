package usecase

type TransactionTopUpRequest struct {
	Amount        float64 `json:"amount"`
	Notes         string  `json:"note"`
	UserId        uint    `json:"userId,omitempty" `
	BankAccountId uint    `json:"bankAccountId"`
}
