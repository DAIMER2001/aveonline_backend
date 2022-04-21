package interactor

import (
	"errors"

	"github.com/DAIMER2001/demo-rest-api/domain/model"
	"github.com/DAIMER2001/demo-rest-api/usecase/presenter"
	"github.com/DAIMER2001/demo-rest-api/usecase/repository"
)

type invoiceInteractor struct {
	InvoiceRepository repository.InvoiceRepository
	InvoicePresenter  presenter.InvoicePresenter
	DBRepository      repository.DBRepository
}

type InvoiceInteractor interface {
	Get(u []*model.Invoice) ([]*model.Invoice, error)
	FindForDates(init string, end string) (u []*model.Invoice, err error)
	Find(u *model.Invoice, id string) (*model.Invoice, error)
	Save(u *model.Invoice) (*model.Invoice, error)
	Delete(id string) (*model.Invoice, error)
	SaveInvoiceMedicines(u *model.InvoiceMedicines) (*model.InvoiceMedicines, error)
}

func NewInvoiceInteractor(r repository.InvoiceRepository, p presenter.InvoicePresenter, d repository.DBRepository) InvoiceInteractor {
	return &invoiceInteractor{r, p, d}
}

func (us *invoiceInteractor) Get(u []*model.Invoice) ([]*model.Invoice, error) {
	u, err := us.InvoiceRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (us *invoiceInteractor) FindForDates(init string, end string) (u []*model.Invoice, err error) {
	u, err = us.InvoiceRepository.FindForDates(init, end)
	// u, err := us.InvoiceRepository.FindForDates(u, dates.DateEnd, dates.DateInit)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (us *invoiceInteractor) Find(u *model.Invoice, id string) (*model.Invoice, error) {
	u, err := us.InvoiceRepository.Find(u, id)
	if err != nil || u.InvoiceID == 0 {
		return nil, err
	}

	return u, nil
}

func (us *invoiceInteractor) Save(u *model.Invoice) (*model.Invoice, error) {
	data, err := us.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		u, err := us.InvoiceRepository.Save(u)

		// do mailing
		// do logging
		// do another process
		return u, err
	})
	invoice, ok := data.(*model.Invoice)

	if !ok {
		return nil, errors.New("cast error")
	}

	if !errors.Is(err, nil) {
		return nil, err
	}

	return invoice, nil
}

func (us *invoiceInteractor) SaveInvoiceMedicines(u *model.InvoiceMedicines) (*model.InvoiceMedicines, error) {
	data, err := us.DBRepository.Transaction(func(interface{}) (interface{}, error) {
		u, err := us.InvoiceRepository.SaveInvoiceMedicines(u)

		// do mailing
		// do logging
		// do another process
		return u, err
	})
	invoice, ok := data.(*model.InvoiceMedicines)

	if !ok {
		return nil, errors.New("cast error")
	}

	if !errors.Is(err, nil) {
		return nil, err
	}

	return invoice, nil
}

func (us *invoiceInteractor) Delete(id string) (u *model.Invoice, err error) {

	if u, err = us.InvoiceRepository.Delete(id); err != nil {
		return nil, err
	}

	return u, nil
}
