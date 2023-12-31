package repositorygorm

import (
	"errors"
	"fmt"
	"synergize/entity"

	"gorm.io/gorm"
)

type BankAccountRepository struct {
	db *gorm.DB
}

func NewBankAccountRepository(db *gorm.DB) *BankAccountRepository {
	return &BankAccountRepository{db: db}
}

func (b *BankAccountRepository) Create(ent *entity.BankAccount) (err error) {
	if err = b.db.Create(ent).Error; err != nil {
		return
	}
	return
}

func (b *BankAccountRepository) FindbyId(bankAccountId int) (ent *entity.BankAccount, err error) {
	query := b.db.Preload("User").Find(&ent, bankAccountId)

	if errors.Is(query.Error, gorm.ErrRecordNotFound) || query.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
		return
	}

	if query.Error != nil {
		err = query.Error
		return
	}

	return
}

func (b *BankAccountRepository) Update(ent *entity.BankAccount) (err error) {
	if err = b.db.Save(&ent).Error; err != nil {
		return
	}
	return
}

func (b *BankAccountRepository) FindBankAccountByUserId(UserId uint) (ent *entity.BankAccount, err error) {
	query := b.db.Preload("User").Find(&ent, "user_id = ?", UserId)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) || query.RowsAffected == 0 {
		errMsg := "Bank account not found"
		err = errors.New(errMsg)
		return
	}

	if query.Error != nil {
		err = query.Error
		return
	}

	return
}

func (b *BankAccountRepository) CheckUserIdExist(UserId uint) (err error) {
	var bankAccount entity.BankAccount
	query := b.db.Find(&bankAccount, "user_id = ?", UserId)

	if query.Error != nil {
		err = query.Error
		return
	}

	if query.RowsAffected > 0 {
		errMsg := fmt.Sprintf("Bank Account already exist %s : %s", bankAccount.BankName, bankAccount.AccountBankNumber)
		err = errors.New(errMsg)
		return
	}

	return
}
