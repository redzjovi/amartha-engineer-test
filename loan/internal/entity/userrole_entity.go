package entity

import "time"

type UserRole struct {
	UserID    uint      `gorm:"column:user_id"`
	RoleCode  RoleCode  `gorm:"column:role_code"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (ur *UserRole) TableName() string {
	return "user_role"
}
