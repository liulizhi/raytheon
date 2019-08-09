package datamodels

import (
	"github.com/jinzhu/gorm"
)

// User user model field
type User struct {
	gorm.Model
	UserName string `gorm:"size:255;UNIQUE;not null;column:username" form:"username" binding:"required" valid:"required,is-uniq"`
	Password string `gorm:"size:255;column:password" form:"password" binding:"required"`
	Phone    string `gorm:"size:11;column:phone"`
	Role     string `gorm:"size:16;column:role"`
	Email    string `gorm:"size:32;column:email"`
	Avatar   string `gorm:"size:1024;column:avatar"`
}

// Tenants tenants information
type Tenants struct {
	gorm.Model
	Name  string `gorm:"size:255;UNIQUE;not null;column:name"`
	Cname string `gorm:"size:255;column:cname"`
}

// UserTenant user and tenant
type UserTenant struct {
	gorm.Model
	TenantID string `gorm:"size:255;UNIQUE;not null;column:tenant_id"`
	UserID   string `gorm:"size:255;UNIQUE;not null;column:user_id"`
}
