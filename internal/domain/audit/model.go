package audit

import (
	"gorm.io/datatypes"
	"time"
)

type AuditLog struct {
	ID           uint           `gorm:"primaryKey;autoIncrement"`
	UserID       uint           `gorm:"not null;index" json:"user_id"`
	Action       string         `gorm:"type:varchar(20);not null" json:"action"`
	ResourceType string         `gorm:"type:varchar(50);not null" json:"resource_type"`
	ResourceID   string         `gorm:"not null;index" json:"resource_id"`
	OldData      datatypes.JSON `gorm:"type:jsonb" json:"old_data"`
	NewData      datatypes.JSON `gorm:"type:jsonb" json:"new_data"`
	IPAddress    string         `gorm:"type:varchar(45)" json:"ip_address"`
	UserAgent    string         `gorm:"type:text" json:"user_agent"`
	Description  string         `gorm:"type:text" json:"description"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
}
