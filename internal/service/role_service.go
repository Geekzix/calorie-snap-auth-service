package service

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/Geekzix/calorie-snap-auth-service/internal/domain"
	"github.com/Geekzix/calorie-snap-auth-service/internal/repository"
	eror "github.com/Geekzix/calorie-snap-commons/common/constant"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

var roleNamePattern = regexp.MustCompile(`^[A-Z][A-Z0-9_]{1,49}$`)

type RoleSvc struct {
	repo repository.RoleRepository
}

func NewRoleSvc(repo *repository.RoleRepository) *RoleSvc {
	return &RoleSvc{repo: *repo}
}

func (s *RoleSvc) Create(ctx context.Context, name string, desc string) (*domain.Role, error) {
	name = strings.ToLower(strings.TrimSpace(name))
	desc = strings.TrimSpace(desc)

	if !roleNamePattern.MatchString(name) || name == "" || len(name) > 50 {
		return nil, eror.InvalidInput("invalid role name")
	}

	role := &domain.Role{
		ID:          uuid.New(),
		Name:        name,
		Description: desc,
	}

	if err := s.repo.Create(ctx, role); err != nil {
		if isDuplicateKey(err) {
			return nil, eror.AlreadyExits("role already exits with this name")
		}
		return nil, err
	}
	return role, nil
}

func (s *RoleSvc) GetByID(ctx context.Context, id uuid.UUID) (*domain.Role, error) {
	role, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, eror.NotFound("role not found")
		}
		return nil, err
	}
	return role, nil
}

func (s *RoleSvc) List(ctx context.Context) ([]domain.Role, error) {
	return s.repo.List(ctx)
}

func (s *RoleSvc) Update(ctx context.Context, id uuid.UUID, name string, desc string) (*domain.Role, error) {
	name = strings.ToLower(strings.TrimSpace(name))
	desc = strings.TrimSpace(desc)

	if !roleNamePattern.MatchString(name) || len(name) > 50 {
		return nil, eror.InvalidInput("invalid role name")
	}

	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, eror.NotFound("role not found")
		}
		return nil, err
	}

	existing.Name = name
	existing.Description = desc
	existing.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, existing); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, eror.NotFound("role not found")
		}
		if isDuplicateKey(err) {
			return nil, eror.AlreadyExits("role already exits with this name")
		}
		return nil, err
	}

	return existing, nil
}

func (s *RoleSvc) Delete(ctx context.Context, id uuid.UUID) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return eror.NotFound("role not found")
		}
		return err
	}
	return nil
}

func isDuplicateKey(err error) bool {
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return true
	}
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}
