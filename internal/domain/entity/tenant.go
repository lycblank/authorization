package entity

// Tenant 聚合根
type Tenant struct {
	TenantId    int64  `gorm:"column:tenant_id;autoIncrement"`
	AppId       string `gorm:"column:app_id;uniqueIndex"`
	AppSecret   string `gorm:"column:app_secret"`
	Disable     bool   `gorm:"column:disable"`
	DisableTime int64  `gorm:"column:disable_time"`
	CreateTime  int64  `gorm:"column:create_time"`
}

func (t Tenant) TableName() string {
	return "tenant"
}
