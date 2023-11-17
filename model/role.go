package model

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Code string `json:"code"`
	Name string `json:"name" gorm:"unique"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
