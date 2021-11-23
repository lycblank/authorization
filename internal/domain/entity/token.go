package entity

// Token 值对象
type Token struct {
	Uid                    int64  `gorm:"column:uid;primaryKey"`
	Token                  string `gorm:"column:token"`
	TokenExpireTime        int64  `gorm:"column:token_expire_time"`
	RefreshToken           string `gorm:"column:refresh_token"`
	RefreshTokenExpireTime int64  `gorm:"column:refresh_token_expire_time"`
	Disable                bool   `gorm:"column:disable"`
	DisableTime            int64  `gorm:"column:disable_time"`
}

func (t Token) TableName() string {
	return "token"
}
