package repositorygorm

import (
	"errors"
	"synergize/entity"

	"gorm.io/gorm"
)

type RepositoryGormTransaction struct {
	db *gorm.DB
}

func NewRepositoryGormTransaction(db *gorm.DB) *RepositoryGormTransaction {
	return &RepositoryGormTransaction{db: db}
}

func (r *RepositoryGormTransaction) CreateTransaction(ent *entity.Transaction) (err error) {
	balance, balanceErr := r.updateBalance(ent)
	if balanceErr != nil {
		err = balanceErr
		return
	}

	err = r.db.Transaction(func(tx *gorm.DB) (err error) {
		// transaction insert
		if err = tx.Create(&ent).Error; err != nil {
			return
		}
		// update balance
		balance.Amount += ent.Amount
		if err = tx.Model(&entity.Balance{}).Where("user_id = ?", balance.UserId).Update("amount", balance.Amount).Error; err != nil {
			return
		}

		return
	})

	if err != nil {
		return
	}

	return
}

func (r *RepositoryGormTransaction) updateBalance(ent *entity.Transaction) (bal *entity.Balance, err error) {
	query := r.db.Where("user_id = ?", ent.UserId).Find(&bal)

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
