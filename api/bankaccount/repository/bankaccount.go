package repository

import "synergize/entity"

type BankAccountRepository interface {
	Create(ent *entity.BankAccount) (err error)
	FindbyId(bankAccountId int) (ent *entity.BankAccount, err error)
	Update(ent *entity.BankAccount) (err error)
	CheckUserIdExist(UserId uint) (err error)
	FindBankAccountByUserId(UserId uint) (ent *entity.BankAccount, err error)
}
