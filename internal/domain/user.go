package domain

import (
	"github.com/Geekzix/calorie-snap-commons/domain/common"
	"github.com/Geekzix/calorie-snap-commons/domain/users/constant"
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID           `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username string              `gorm:"type:varchar(255);not null"`
	Email    string              `gorm:"type:varchar(100);not null;uniqueIndex:ux_user_email"`
	Phone    string              `gorm:"type:varchar(15);not null;uniqueIndex:ux_user_phone"`
	Password string              `gorm:"type:varchar(255);not null"`
	Gender   constant.Gender     `gorm:"type:int;not null"`
	Status   constant.UserStatus `gorm:"type:varchar(32);not null;default:PENDING;index:idx_users_status"`

	Role Role `gorm:"foreignKey:RoleID;references:ID:constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	common.Audit
}
