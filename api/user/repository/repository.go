package repository

import (
	usecaseuser "synergize/api/user/usecase"
	"synergize/entity"
)

type RepositoryUser interface {
	UserList(cmd usecaseuser.UserQueryParam) (ent []*entity.UserListWithBankAccount, err error)
	UserDetail(userId uint) (ent *entity.User, err error)
}
