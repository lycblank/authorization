package entity

// Account 聚合根
type Account struct {
	Uid           int64  `gorm:"column:uid;autoIncrement"`
	UserName      string `gorm:"column:user_name;uniqueIndex:ut_unique_index"`
	Password      string `gorm:"column:password"`
	Salt          string `gorm:"column:salt"`
	CreateTime    int64  `gorm:"column:create_time"`
	LastLoginTime int64  `gorm:"column:last_login_time"`
	Disable       bool   `gorm:"column:disable"`
	DisableTime   int64  `gorm:"column:disable_time"`
	TenantId      int64  `gorm:"column:tenant_id;uniqueIndex:ut_unique_index"`
	Token         Token  `gorm:"foreignkey:Uid"` // 1对1
}

func (a Account) TableName() string {
	return "account"
}
