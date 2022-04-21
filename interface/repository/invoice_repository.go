package repository

import (
	"errors"
	"fmt"

	"github.com/DAIMER2001/demo-rest-api/domain/model"
	"github.com/DAIMER2001/demo-rest-api/usecase/repository"
	"gorm.io/gorm"
)

type invoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) repository.InvoiceRepository {
	return &invoiceRepository{db}
}

func (ur *invoiceRepository) FindAll(u []*model.Invoice) ([]*model.Invoice, error) {
	err := ur.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *invoiceRepository) FindForDates(init string, end string) (u []*model.Invoice, err error) {
	// TimeInit, _ := time.Parse(time.RFC3339, init)
	// TimeEnd, _ := time.Parse(time.RFC3339, end)
	fmt.Println(init)
	fmt.Println(end)
	err = ur.db.Where("date_payment >= ? AND date_payment <= ?", init, end).Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *invoiceRepository) Find(u *model.Invoice, id string) (*model.Invoice, error) {

	if err := ur.db.Find(u, id).Error; err != nil || u.InvoiceID == 0 {
		return nil, err
	}
	return u, nil
}

func (ur *invoiceRepository) Save(u *model.Invoice) (*model.Invoice, error) {

	if err := ur.db.Save(u).Error; !errors.Is(err, nil) {
		return nil, err
	}
	return u, nil
}
func (ur *invoiceRepository) SaveInvoiceMedicines(u *model.InvoiceMedicines) (*model.InvoiceMedicines, error) {

	if err := ur.db.Save(u).Error; !errors.Is(err, nil) {
		return nil, err
	}
	return u, nil
}
func (ur *invoiceRepository) Delete(id string) (u *model.Invoice, err error) {

	if err := ur.db.Find(u, id).Error; !errors.Is(err, nil) || u == nil {
		return nil, err
	}

	if err := ur.db.Delete(u).Error; !errors.Is(err, nil) {
		return nil, err
	}

	return u, nil
}
