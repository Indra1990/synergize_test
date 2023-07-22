package repositorygorm

import (
	"errors"
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

	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		err = query.Error
		return
	}

	if query.Error != nil {
		err = query.Error
		return
	}

	if query.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
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
