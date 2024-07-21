package repository

import (
	"loan/internal/entity"

	"gorm.io/gorm"
)

type UserRoleRepository struct {
	Repository[entity.UserRole]
}

func NewUserRoleRepository() *UserRoleRepository {
	return &UserRoleRepository{}
}

func (r *UserRoleRepository) CountByUserIdAndRoleCode(db *gorm.DB, userId uint, roleCode entity.RoleCode) (total int64, err error) {
	err = db.
		Model(entity.UserRole{}).
		Where("user_id = ?", userId).
		Where("role_code = ?", roleCode).
		Count(&total).Error
	return total, err
}
