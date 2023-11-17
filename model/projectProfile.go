package model

import (
	"time"

	"gorm.io/gorm"
)

type ProjectProfile struct {
	Id        uint `json:"id" gorm:"primaryKey"`
	ProfileId uint `json:"profileId"`
	ProjectId uint `json:"projectId"`

	Profile *Profile `json:"profile" gorm:"foreignKey:ProfileId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Project *Project `json:"project" gorm:"foreignKey:ProjectId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
