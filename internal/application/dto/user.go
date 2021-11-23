package dto

type RegisterAccountCmd struct {
	UserName string `validate:"required"`
	Password string `validate:"required"`
	AppId    string `validate:"required"`
}
