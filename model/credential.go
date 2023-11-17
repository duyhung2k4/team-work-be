package model

import (
	"time"

	"gorm.io/gorm"
)

type Credential struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	RoleId   uint   `json:"roleId"`
	Email    string `json:"email" gorm:"unique"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-"`

	Role *Role `json:"role" gorm:"foreignKey:RoleId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
