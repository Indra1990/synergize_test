package usecase

type BankAccountRequest struct {
	BankName          string `json:"bankName"`
	AccountName       string `json:"accountName"`
	AccountBankNumber string `json:"accountBankNumber"`
	UserId            uint
}

type BankAccountRequestUpdate struct {
	ID                uint
	BankName          string `json:"bankName"`
	AccountName       string `json:"accountName"`
	AccountBankNumber string `json:"accountBankNumber"`
	UserId            uint
}

type BankAccountResponse struct {
	UserId                  uint                     `json:"userId,omitempty"`
	BankName                string                   `json:"bankName"`
	AccountName             string                   `json:"accountName"`
	AccountBankNumber       string                   `json:"accountBankNumber"`
	CreatedAt               string                   `json:"createdAt,omitempty"`
	UpdatedAt               string                   `json:"updatedAt,omitempty"`
	BankAccountUserResponse *BankAccountUserResponse `json:"user,omitempty"`
}

type BankAccountUserResponse struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}
