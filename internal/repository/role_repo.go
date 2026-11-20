package repository

import (
	"context"

	"github.com/Geekzix/calorie-snap-auth-service/internal/domain"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepo(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (r *RoleRepository) Create(ctx context.Context, role *domain.Role) error {
	return gorm.G[domain.Role](r.db).Create(ctx, role)
}
