package repository

import "github.com/lycblank/authorization/internal/domain/entity"

type AccountRepository interface {
	Save(account *entity.Account) error
	Get(uid int64) (*entity.Account, error)
}
