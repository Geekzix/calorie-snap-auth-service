package service

import (
	"context"
	"errors"
	"regexp"
	"strings"

	"github.com/Geekzix/calorie-snap-auth-service/internal/domain"
	"github.com/Geekzix/calorie-snap-auth-service/internal/repository"
	eror "github.com/Geekzix/calorie-snap-commons/common/constant"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

var roleCodePattern = regexp.MustCompile(`^[A-Z][A-Z0-9_]{1,49}$`)

type RoleSvc struct {
	repo repository.RoleRepository
}

func NewRoleSvc(repo *repository.RoleRepository) *RoleSvc {
	return &RoleSvc{repo: *repo}
}

func (s *RoleSvc) Create(ctx context.Context, name string, desc string) (*domain.Role, error) {
	name = strings.ToLower(strings.TrimSpace(name))
	desc = strings.TrimSpace(desc)

	if !roleCodePattern.MatchString(name) || name == "" || len(name) > 50 {
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

func isDuplicateKey(err error) bool {
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return true
	}
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}
