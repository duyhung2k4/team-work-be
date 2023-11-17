package model

import (
	"time"

	"gorm.io/gorm"
)

type Team struct {
	Id        uint     `json:"id" gorm:"primaryKey"`
	CreaterId uint     `json:"createrId"`
	Code      string   `json:"code"`
	Name      string   `json:"name"`
	Password  string   `json:"password"`
	Creater   *Profile `json:"creater" gorm:"foreignKey:CreaterId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
