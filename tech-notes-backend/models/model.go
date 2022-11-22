package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time      `json:"created_at"`
	CreatedBy string         `json:"created_by" gorm:"type:varchar(50);DEFAULT:NULL"`
	UpdatedAt time.Time      `json:"updated_at"`
	UpdatedBy string         `json:"updated_by" gorm:"type:varchar(50);DEFAULT:NULL"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	DeletedBy string         `json:"-" gorm:"type:varchar(50);DEFAULT:NULL"`
}
