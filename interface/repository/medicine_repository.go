package repository

import (
	"errors"

	"github.com/DAIMER2001/demo-rest-api/domain/model"
	"github.com/DAIMER2001/demo-rest-api/usecase/repository"
	"gorm.io/gorm"
)

type medicineRepository struct {
	db *gorm.DB
}

func NewMedicineRepository(db *gorm.DB) repository.MedicineRepository {
	return &medicineRepository{db}
}

func (ur *medicineRepository) FindAll(u []*model.Medicine) ([]*model.Medicine, error) {
	// err := ur.db.Preload("Promotions").Find(&u).Error
	// err := ur.db.Preload("Promotions", "start_date ").Order(`"Promotions"."start_date" ASC`).Find(&u).Error
	// err := ur.db.Preload("Promotions", "order(`promotions`.start_date ASC`)").Find(&u).Error

	err := ur.db.Preload("Promotions", func(db *gorm.DB) *gorm.DB {
		return db.Order("promotion.start_date ASC")
	}).Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *medicineRepository) Find(u *model.Medicine, id string) (*model.Medicine, error) {

	if err := ur.db.Find(u, id).Error; err != nil || u.MedicineID == 0 {
		return nil, err
	}
	return u, nil
}

func (ur *medicineRepository) Save(u *model.Medicine) (*model.Medicine, error) {

	if err := ur.db.Save(u).Error; !errors.Is(err, nil) {
		return nil, err
	}
	return u, nil
}

func (ur *medicineRepository) Delete(id string) (u *model.Medicine, err error) {

	if err := ur.db.Find(u, id).Error; !errors.Is(err, nil) || u == nil {
		return nil, err
	}

	if err := ur.db.Delete(u).Error; !errors.Is(err, nil) {
		return nil, err
	}

	return u, nil
}
