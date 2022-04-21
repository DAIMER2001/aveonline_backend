package interactor

import (
	"errors"

	"github.com/DAIMER2001/demo-rest-api/domain/model"
	"github.com/DAIMER2001/demo-rest-api/usecase/presenter"
	"github.com/DAIMER2001/demo-rest-api/usecase/repository"
)

type promotionInteractor struct {
	PromotionRepository repository.PromotionRepository
	PromotionPresenter  presenter.PromotionPresenter
	DBRepository        repository.DBRepository
}

type PromotionInteractor interface {
	Get(u []*model.Promotion) ([]*model.Promotion, error)
	Find(u *model.Promotion, id string) (*model.Promotion, error)
	Save(u *model.Promotion) (*model.Promotion, error)
	Delete(id string) (*model.Promotion, error)
	SavePromotionMedicines(u *model.PromotionMedicines) (*model.PromotionMedicines, error)
}

func NewPromotionInteractor(r repository.PromotionRepository, p presenter.PromotionPresenter, d repository.DBRepository) PromotionInteractor {
	return &promotionInteractor{r, p, d}
}

func (us *promotionInteractor) Get(u []*model.Promotion) ([]*model.Promotion, error) {
	u, err := us.PromotionRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (us *promotionInteractor) Find(u *model.Promotion, id string) (*model.Promotion, error) {
	u, err := us.PromotionRepository.Find(u, id)
	if err != nil || u.PromotionID == 0 {
		return nil, err
	}

	return u, nil
}

func (us *promotionInteractor) Save(u *model.Promotion) (*model.Promotion, error) {
	data, err := us.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		u, err := us.PromotionRepository.Save(u)

		// do mailing
		// do logging
		// do another process
		return u, err
	})
	promotion, ok := data.(*model.Promotion)

	if !ok {
		return nil, errors.New("cast error")
	}

	if !errors.Is(err, nil) {
		return nil, err
	}

	return promotion, nil
}

func (us *promotionInteractor) Delete(id string) (u *model.Promotion, err error) {

	if u, err = us.PromotionRepository.Delete(id); err != nil {
		return nil, err
	}

	return u, nil
}

func (us *promotionInteractor) SavePromotionMedicines(u *model.PromotionMedicines) (*model.PromotionMedicines, error) {

	data, err := us.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		u, err := us.PromotionRepository.SavePromotionMedicines(u)

		// do mailing
		// do logging
		// do another process
		return u, err
	})
	promotion, ok := data.(*model.PromotionMedicines)

	if !ok {
		return nil, errors.New("cast error")
	}

	if !errors.Is(err, nil) {
		return nil, err
	}

	return promotion, nil
}
