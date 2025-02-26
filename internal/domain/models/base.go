package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type BaseModel struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt time.Time
	IsDeleted soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"`
}
