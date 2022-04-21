package repository

import "github.com/DAIMER2001/demo-rest-api/domain/model"

type PromotionRepository interface {
	FindAll(u []*model.Promotion) ([]*model.Promotion, error)
	Find(u *model.Promotion, id string) (*model.Promotion, error)
	Save(u *model.Promotion) (*model.Promotion, error)
	Delete(id string) (*model.Promotion, error)
	SavePromotionMedicines(u *model.PromotionMedicines) (*model.PromotionMedicines, error)
}
