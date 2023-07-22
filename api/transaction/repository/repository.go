package repository

import "synergize/entity"

type RepositoryTransaction interface {
	CreateTransaction(ent *entity.Transaction) (err error)
}
