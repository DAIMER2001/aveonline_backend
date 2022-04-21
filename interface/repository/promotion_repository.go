package repository

import (
	"errors"
	"fmt"

	"github.com/DAIMER2001/demo-rest-api/domain/model"
	"github.com/DAIMER2001/demo-rest-api/usecase/repository"
	"gorm.io/gorm"
)

type promotionRepository struct {
	db *gorm.DB
}

func NewPromotionRepository(db *gorm.DB) repository.PromotionRepository {
	return &promotionRepository{db}
}

func (ur *promotionRepository) FindAll(u []*model.Promotion) ([]*model.Promotion, error) {
	// err := ur.db.Find(&u).Error

	err := ur.db.Preload("Medicines").Find(&u).Error

	fmt.Println(u)
	fmt.Println(&u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *promotionRepository) Find(u *model.Promotion, id string) (*model.Promotion, error) {

	if err := ur.db.Find(u, id).Error; err != nil || u.PromotionID == 0 {
		return nil, err
	}
	return u, nil
}

func (ur *promotionRepository) Save(u *model.Promotion) (*model.Promotion, error) {

	if err := ur.db.Omit("Medicines").Save(u).Error; !errors.Is(err, nil) {
		return nil, err
	}
	return u, nil
}

func (ur *promotionRepository) Delete(id string) (u *model.Promotion, err error) {

	if err := ur.db.Find(u, id).Error; !errors.Is(err, nil) || u == nil {
		return nil, err
	}

	if err := ur.db.Delete(u).Error; !errors.Is(err, nil) {
		return nil, err
	}

	return u, nil
}

func (ur *promotionRepository) SavePromotionMedicines(u *model.PromotionMedicines) (*model.PromotionMedicines, error) {

	if err := ur.db.Save(u).Error; !errors.Is(err, nil) {
		return nil, err
	}
	return u, nil
}
