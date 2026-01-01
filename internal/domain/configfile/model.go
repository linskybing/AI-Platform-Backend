package configfile

import "time"

type ConfigFile struct {
	CFID      uint      `gorm:"primaryKey;column:cf_id"`
	Filename  string    `gorm:"size:200;not null"`
	Content   string    `gorm:"size:5000"`
	ProjectID uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"column:create_at"`
}
