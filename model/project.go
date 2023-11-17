package model

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	CreaterId uint   `json:"createrId"`
	TeamId    uint   `json:"teamId"`
	Name      string `json:"name"`

	Creater *Profile `json:"creater" gorm:"foreignKey:CreaterId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Team    *Team    `json:"team" gorm:"foreignKey:TeamId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
