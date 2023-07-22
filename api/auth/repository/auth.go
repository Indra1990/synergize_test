package repository

import "synergize/entity"

type AuthRepository interface {
	CreateUser(ent *entity.User) (err error)
	FindByUserEmail(email string) (ent *entity.User, err error)
}
