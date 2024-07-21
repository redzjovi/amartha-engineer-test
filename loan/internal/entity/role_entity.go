package entity

import "time"

type RoleCode string

const (
	RoleCodeAdmin = "admin"
)

type Role struct {
	Code      RoleCode  `gorm:"column:code;primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (r *Role) TableName() string {
	return "roles"
}
