package service

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"time"

	"github.com/lycblank/authorization/internal/domain/repository"

	"github.com/lycblank/authorization/internal/domain/entity"
)

type AccountService struct {
	passwordWithSaltStrategy repository.PasswordWithSaltStrategy
	accountRepositoryFactory repository.AccountRepositoryFactory
}

func (as *AccountService) RegisterAccount(ctx context.Context, username, password string, tenantId int64) (int64, error) {
	account := new(entity.Account)
	account.UserName = username
	account.Salt = as.genSalt()
	account.Password = as.passwordWithSaltStrategy.AddSalt(password, account.Salt)
	account.TenantId = tenantId

	accountRepository := as.accountRepositoryFactory.CreateAccountRepository()
	accountRepository.Save(account)
}

func (as *AccountService) genSalt() string {
	salt := make([]byte, 8)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		salt = []byte(fmt.Sprintf("%d", time.Now().UnixNano()))
	}
	return string(salt)
}
