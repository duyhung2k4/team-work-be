package model

import (
	"time"

	"gorm.io/gorm"
)

type ApiRoleTeam struct {
	Id         uint `json:"id"`
	RoleTeamId uint `json:"roleTeamId"`
	ApiId      uint `json:"apiId"`
	Enable     bool `json:"enable"`

	RoleTeam *RoleTeam `json:"roleTeam" gorm:"foreignKey:RoleTeamId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Api      *Api      `json:"api" gorm:"foreignKey:ApiId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
