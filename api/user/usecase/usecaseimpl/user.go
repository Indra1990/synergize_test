package usecaseimpl

import (
	"strconv"
	usecasedtobankaccount "synergize/api/bankaccount/usecase"
	repositoryuser "synergize/api/user/repository"
	usecaseuser "synergize/api/user/usecase"
)

type ServiceUser struct {
	repositoryUser repositoryuser.RepositoryUser
}

func NewServiceUser(repositoryUser repositoryuser.RepositoryUser) *ServiceUser {
	return &ServiceUser{
		repositoryUser: repositoryUser,
	}
}

func (s *ServiceUser) UserList(cmd usecaseuser.UserQueryParam) (dto []*usecaseuser.UserListResponse, err error) {
	userList, userListErr := s.repositoryUser.UserList(cmd)
	if userListErr != nil {
		err = userListErr
		return
	}

	if len(userList) == 0 {
		return
	}

	dto = make([]*usecaseuser.UserListResponse, len(userList))
	for idx, usr := range userList {
		user := usecaseuser.UserListResponse{
			ID:          usr.ID,
			Username:    usr.Username,
			Email:       usr.Email,
			PhoneNumber: usr.PhoneNumber,
			Balance:     usr.Balance,
			CreateAt:    usr.CreatedAtUser,
			UpdatedAt:   usr.UpdatedAtUser,
			BankAccount: &usecasedtobankaccount.BankAccountResponse{
				BankName:          usr.BankName,
				AccountName:       usr.AccountName,
				AccountBankNumber: usr.AccountBankNumber,
				CreatedAt:         usr.CreatedAtBankAccount,
				UpdatedAt:         usr.UpdatedAtBankAccount,
			},
		}
		dto[idx] = &user
	}

	return
}

func (s *ServiceUser) UserDetail(userId uint) (dto *usecaseuser.UserListResponse, err error) {
	var balance string
	var bankAccRes usecasedtobankaccount.BankAccountResponse

	userDetail, userDetailErr := s.repositoryUser.UserDetail(userId)
	if userDetailErr != nil {
		err = userDetailErr
		return
	}

	if userDetail.BankAccount != nil {
		bankAccRes.BankName = userDetail.BankAccount.BankName
		bankAccRes.AccountName = userDetail.BankAccount.AccountName
		bankAccRes.AccountBankNumber = userDetail.BankAccount.AccountBankNumber

	}

	if userDetail.Balance != nil {
		balance = strconv.FormatFloat(userDetail.Balance.Amount, 'f', -1, 64)
	}

	detailUser := &usecaseuser.UserListResponse{
		ID:          userDetail.ID,
		Username:    userDetail.Username,
		Email:       userDetail.Email,
		PhoneNumber: userDetail.PhoneNumber,
		Balance:     balance,
		BankAccount: &bankAccRes,
	}

	dto = detailUser
	return
}
