package repositorygorm

import (
	"errors"
	"fmt"
	"strings"
	usecaseuser "synergize/api/user/usecase"
	"synergize/entity"

	"gorm.io/gorm"
)

type RepositoryGormUser struct {
	db *gorm.DB
}

func NewRepositoryGormUser(db *gorm.DB) *RepositoryGormUser {
	return &RepositoryGormUser{
		db: db,
	}
}

func (r *RepositoryGormUser) UserList(cmd usecaseuser.UserQueryParam) (ent []*entity.UserListWithBankAccount, err error) {

	query := r.db.Model(&entity.User{}).
		Select(
			"users.id",
			"users.username",
			"users.email",
			"users.phone_number",
			"TO_CHAR(users.created_at,'DD Mon YYYY HH24:MI:SS') as created_at_user",
			"TO_CHAR(users.updated_at,'DD Mon YYYY HH24:MI:SS') as updated_at_user",
			"bank_accounts.bank_name",
			"bank_accounts.account_name",
			"bank_accounts.account_bank_number",
			"TO_CHAR(bank_accounts.created_at,'DD Mon YYYY HH24:MI:SS') as created_at_bank_account",
			"TO_CHAR(bank_accounts.updated_at,'DD Mon YYYY HH24:MI:SS') as updated_at_bank_account",
			"balances.amount as balance",
		).
		Joins("left join bank_accounts on users.id = bank_accounts.user_id").
		Joins("left join balances on users.id = balances.user_id")

	if cmd.Username != "" {
		query.Where("users.username = ?", cmd.Username)
	}

	if cmd.AccountName != "" {
		query.Where("bank_accounts.account_name = ?", cmd.AccountName)
	}

	if cmd.AccountNumber != "" {
		query.Where("bank_accounts.account_bank_number = ?", cmd.AccountNumber)
	}

	if cmd.AccountBankName != "" {
		query.Where("bank_accounts.bank_name = ?", strings.ToUpper(cmd.AccountBankName))
	}

	if cmd.BalanceStart > 0 {
		query.Where("balances.amount >= ?", cmd.BalanceStart)
	}

	if cmd.BalanceEnd > 0 {
		query.Where("balances.amount <= ?", cmd.BalanceEnd)
	}

	if cmd.RegisterAt != "" {
		start := fmt.Sprintf("%s 00:00:00", cmd.RegisterAt)
		end := fmt.Sprintf("%s 23:59:00", cmd.RegisterAt)
		query.Where("users.created_at >= ?", start).Where("users.created_at <= ?", end)
	}

	query.Order("id desc").
		Find(&ent)

	if query.Error != nil {
		err = query.Error
		return
	}

	return
}

func (r *RepositoryGormUser) UserDetail(userId uint) (ent *entity.User, err error) {
	query := r.db.Preload("BankAccount").Preload("Balance").Find(&ent, userId)

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
