package domain

import (
	"github.com/Geekzix/calorie-snap-commons/domain/common"
	"github.com/google/uuid"
)

type Permission struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Code        string    `gorm:"type:varchar(100);not null;uniqueIndex:ux_permissions_code"`
	Description string    `gorm:"type:text"`

	common.Audit
}
