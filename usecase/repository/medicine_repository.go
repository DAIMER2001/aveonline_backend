package repository

import "github.com/DAIMER2001/demo-rest-api/domain/model"

type MedicineRepository interface {
	FindAll(u []*model.Medicine) ([]*model.Medicine, error)
	Find(u *model.Medicine, id string) (*model.Medicine, error)
	Save(u *model.Medicine) (*model.Medicine, error)
	Delete(id string) (*model.Medicine, error)
}
