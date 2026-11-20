package repository

import (
	"context"
	"errors"

	"github.com/Geekzix/calorie-snap-auth-service/internal/domain"
	"github.com/google/uuid"
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

func (r *RoleRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Role, error) {
	role, err := gorm.G[domain.Role](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepository) List(ctx context.Context) ([]domain.Role, error) {
	return gorm.G[domain.Role](r.db).Order("created_at ASC").Find(ctx)
}

func (r *RoleRepository) Update(ctx context.Context, role *domain.Role) error {
	rows, err := gorm.G[domain.Role](r.db).
		Where("id = ?", role.ID).
		Updates(ctx, domain.Role{
			Name:        role.Name,
			Description: role.Description,
			UpdatedBy:   role.UpdatedBy,
		})
	if err != nil {
		return err
	}
	if rows == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *RoleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	rows, err := gorm.G[domain.Role](r.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}
	if rows == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
