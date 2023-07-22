package repositorygorm

import (
	"errors"
	"synergize/entity"

	"gorm.io/gorm"
)

type RepositoryGorm struct {
	db *gorm.DB
}

func NewRepositoryGorm(db *gorm.DB) *RepositoryGorm {
	return &RepositoryGorm{
		db: db,
	}
}

func (r *RepositoryGorm) CreateUser(ent *entity.User) (err error) {
	if err = r.db.Create(&ent).Error; err != nil {
		return
	}

	err = r.db.Transaction(func(tx *gorm.DB) (err error) {
		balance := entity.Balance{
			Amount: 0,
			UserId: ent.ID,
		}

		if err = tx.Create(&balance).Error; err != nil {
			return
		}

		return
	})

	return
}

func (r *RepositoryGorm) FindByUserEmail(email string) (ent *entity.User, err error) {
	if err = r.db.Where("email = ?", email).Find(&ent).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
		return
	}

	return
}
