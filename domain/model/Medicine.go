package model

import (
	"time"
)

type Medicine struct {
	MedicineID uint    `json:"id" gorm:"primary_key; auto_increment"`
	Name       string  `json:"name" gorm:"type:varchar(50);" validate:"required,min=3,max=50"`
	Price      float64 `json:"price" gorm:"type:decimal(13, 2);" validate:"required,numeric,min=0.1,max=10000000000"`
	Location   string  `json:"location" gorm:"type:varchar(50);" validate:"required,min=3,max=50"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time  `sql:"index"`
	Promotions []Promotion `json:"promotion_medicines" gorm:"many2many:promotion_medicines;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Medicine) TableName() string { return "medicines" }
