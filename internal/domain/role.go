package domain

import (
	"github.com/Geekzix/calorie-snap-commons/domain/common"
	"github.com/google/uuid"
)

type Role struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string    `gorm:"type:varchar(50);not null;uniqueIndex;ux_role_name"`
	Description string    `gorm:"type:varchar(255)"`

	Permission []Permission `gorm:"many2many:role_permission;"`
	common.Audit
}
