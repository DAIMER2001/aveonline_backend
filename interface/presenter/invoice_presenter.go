package presenter

import (
	"github.com/DAIMER2001/demo-rest-api/domain/model"
	"github.com/DAIMER2001/demo-rest-api/usecase/presenter"
)

type invoicePresenter struct{}

func NewInvoicePresenter() presenter.InvoicePresenter {
	return &invoicePresenter{}
}

func (up *invoicePresenter) ResponseInvoices(us []*model.Invoice) []*model.Invoice {
	return us
}
