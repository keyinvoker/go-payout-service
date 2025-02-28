package models

import "time"

type BaseModel struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null"`
	IsDeleted bool      `gorm:"type:boolean;default:false"`
	DeletedAt *time.Time
}
