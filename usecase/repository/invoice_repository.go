package repository

import (
	"github.com/DAIMER2001/demo-rest-api/domain/model"
)

type InvoiceRepository interface {
	FindAll(u []*model.Invoice) ([]*model.Invoice, error)
	FindForDates(init string, end string) (u []*model.Invoice, err error)
	Find(u *model.Invoice, id string) (*model.Invoice, error)
	Save(u *model.Invoice) (*model.Invoice, error)
	Delete(id string) (*model.Invoice, error)
	SaveInvoiceMedicines(u *model.InvoiceMedicines) (*model.InvoiceMedicines, error)
}
