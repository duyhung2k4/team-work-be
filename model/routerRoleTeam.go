package model

import (
	"time"

	"gorm.io/gorm"
)

type RouterRoleTeam struct {
	Id         uint `json:"id"`
	RoleTeamId uint `json:"roleTeamId"`
	RouterId   uint `json:"routerId"`
	Enable     bool `json:"enable"`

	RoleTeam *RoleTeam `json:"roleTeam" gorm:"foreignKey:RoleTeamId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Router   *Router   `json:"router" gorm:"foreignKey:RouterId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
