package repository

import "github.com/lycblank/authorization/internal/domain/entity"

type TenantRepository interface {
	Save(tenant *entity.Tenant)
	Get(tid int64) (*entity.Tenant, error)
}
