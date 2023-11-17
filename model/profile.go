package model

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	Id           uint   `json:"id" gorm:"primaryKey"`
	CredentialId uint   `json:"credentialId" gorm:"unique"`
	Phone        string `json:"phone"`

	Credential *Credential `json:"credential" gorm:"foreignKey:CredentialId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
