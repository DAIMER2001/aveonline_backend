package model

import (
	"time"
)

//swagger:model Promotion
type Promotion struct {
	PromotionID uint       `json:"id" gorm:"primary_key; auto_increment"`
	Description string     `json:"description" gorm:"type:varchar(100);" validate:"required,min=2,max=100"`
	Percentage  float64    `json:"percentage" gorm:"type:decimal(5, 2);" validate:"required,numeric,max=70,min=0.1"`
	StartDate   *time.Time `json:"start_date" validate:"required"`
	EndDate     *time.Time `json:"end_date" validate:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
	Medicines   []Medicine `json:"promotion_medicines" gorm:"many2many:promotion_medicines;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type PromotionMedicines struct {
	PromotionID int `gorm:"primaryKey"`
	MedicineID  int `gorm:"primaryKey"`
}

func (Promotion) TableName() string { return "promotions" }
