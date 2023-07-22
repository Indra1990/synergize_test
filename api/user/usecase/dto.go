package usecase

import (
	usecasedtobankaccount "synergize/api/bankaccount/usecase"
)

type UserListResponse struct {
	ID          uint                                       `json:"userId"`
	Username    string                                     `json:"username"`
	Email       string                                     `json:"email"`
	PhoneNumber string                                     `json:"phoneNumber"`
	Balance     string                                     `json:"balance"`
	CreateAt    string                                     `json:"createdAt,omitempty"`
	UpdatedAt   string                                     `json:"updatedAt,omitempty"`
	BankAccount *usecasedtobankaccount.BankAccountResponse `json:"bankAccount,omitempty"`
}

type UserQueryParam struct {
	Username        string
	AccountName     string
	AccountNumber   string
	AccountBankName string
	RegisterAt      string
	BalanceStart    float64
	BalanceEnd      float64
}
