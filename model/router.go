package model

import (
	"time"

	"gorm.io/gorm"
)

type Router struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Path string `json:"path" gorm:"unique"`
	Name string `json:"name"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
