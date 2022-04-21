package model

import (
	"time"
)

type Invoice struct {
	InvoiceID   uint       `json:"id" gorm:"primary_key; auto_increment" form:"check_in"`
	DatePayment *time.Time `json:"date_payment" validate:"required"`
	FullPayment float64    `json:"full_payment" gorm:"type:decimal(13, 2);" validate:"required,numeric,min=0.1,max=10000000000"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time  `sql:"index"`
	Promotions  []Promotion `json:"promotions" gorm:"many2many:invoice_promotions;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Medicines   []Medicine  `json:"medicines" gorm:"many2many:invoice_medicines;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type InvoiceMedicines struct {
	InvoiceID  int `gorm:"primaryKey"`
	MedicineID int `gorm:"primaryKey"`
}

func (Invoice) TableName() string { return "invoices" }
