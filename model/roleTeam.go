package model

import (
	"time"

	"gorm.io/gorm"
)

type RoleTeam struct {
	Id     uint   `json:"id" gorm:"primaryKey"`
	TeamId uint   `json:"teamId"`
	Name   string `json:"name"`
	Code   string `json:"code"`

	Team *Team `json:"team" gorm:"foreignKey:TeamId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
