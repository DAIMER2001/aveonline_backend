package registry

import (
	"github.com/DAIMER2001/demo-rest-api/interface/controller"
	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Promotion: r.NewPromotionController(),
		Invoice:   r.NewInvoiceController(),
		Medicine:  r.NewMedicineController(),
	}
}
