package model

import (
	"time"

	"gorm.io/gorm"
)

// Người trong 1 team
type TeamProfile struct {
	Id        uint `json:"id" gorm:"primaryKey"`
	TeamId    uint `json:"teamId"`
	ProfileId uint `json:"profileId"`

	Profile *Profile `json:"profile" gorm:"foreignKey:ProfileId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Team    *Team    `json:"team" gorm:"foreignKey:TeamId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
