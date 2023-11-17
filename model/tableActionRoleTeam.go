package model

import (
	"time"

	"gorm.io/gorm"
)

type TableActionRoleTeam struct {
	Id            uint `json:"id" gorm:"primaryKey"`
	RoleTeamId    uint `json:"roleTeamId"`
	TableActionId uint `json:"tableActionId"`
	Create        bool `json:"create"`
	Read          bool `json:"read"`
	Update        bool `json:"update"`
	Delete        bool `json:"delete"`

	RoleTeam    *RoleTeam    `json:"roleTeam" gorm:"foreignKey:RoleTeamId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	TableAction *TableAction `json:"tableAction" gorm:"foreignKey:TableActionId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
