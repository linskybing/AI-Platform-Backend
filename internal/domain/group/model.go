package group

import "time"

type Group struct {
	GID         uint      `gorm:"primaryKey;column:g_id"`
	GroupName   string    `gorm:"size:100;not null"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"column:create_at"`
	UpdatedAt   time.Time `gorm:"column:update_at"`
}

func (Group) TableName() string {
	return "group_list"
}

type UserGroup struct {
	UID       uint      `gorm:"primaryKey;column:u_id"`
	GID       uint      `gorm:"primaryKey;column:g_id"`
	Role      string    `gorm:"type:user_role;default:user;not null"` // ENUM
	CreatedAt time.Time `gorm:"column:create_at"`
	UpdatedAt time.Time `gorm:"column:update_at"`
}

func (UserGroup) TableName() string {
	return "user_group"
}
