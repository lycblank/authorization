package repository

type AccountRepositoryFactory interface {
	CreateAccountRepository() AccountRepository
}
