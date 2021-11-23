package repository

type PasswordWithSaltStrategy interface {
	AddSalt(password string, salt string) string
}
