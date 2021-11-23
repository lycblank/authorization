package application

import (
	"github.com/go-playground/validator/v10"
	"github.com/lycblank/authorization/internal/application/dto"
)

type AccountApp struct {
	structValidate *validator.Validate
}

func (aa *AccountApp) RegisterAccount(cmd dto.RegisterAccountCmd) (int64, error) {
	if err := aa.structValidate.Struct(cmd); err != nil {
		return 0, err
	}

	return 0, nil
}
