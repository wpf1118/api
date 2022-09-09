package entity

import "gorm.io/gorm"

type Base struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt int64          `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64          `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
