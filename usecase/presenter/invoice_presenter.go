package presenter

import "github.com/DAIMER2001/demo-rest-api/domain/model"

type InvoicePresenter interface {
	ResponseInvoices(u []*model.Invoice) []*model.Invoice
}
