package interactor

import (
	"errors"

	"github.com/DAIMER2001/demo-rest-api/domain/model"
	"github.com/DAIMER2001/demo-rest-api/usecase/presenter"
	"github.com/DAIMER2001/demo-rest-api/usecase/repository"
)

type medicineInteractor struct {
	MedicineRepository repository.MedicineRepository
	MedicinePresenter  presenter.MedicinePresenter
	DBRepository       repository.DBRepository
}

type MedicineInteractor interface {
	Get(u []*model.Medicine) ([]*model.Medicine, error)
	Find(u *model.Medicine, id string) (*model.Medicine, error)
	Save(u *model.Medicine) (*model.Medicine, error)
	Delete(id string) (*model.Medicine, error)
}

func NewMedicineInteractor(r repository.MedicineRepository, p presenter.MedicinePresenter, d repository.DBRepository) MedicineInteractor {
	return &medicineInteractor{r, p, d}
}

func (us *medicineInteractor) Get(u []*model.Medicine) ([]*model.Medicine, error) {
	u, err := us.MedicineRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (us *medicineInteractor) Find(u *model.Medicine, id string) (*model.Medicine, error) {
	u, err := us.MedicineRepository.Find(u, id)
	if err != nil || u.MedicineID == 0 {
		return nil, err
	}

	return u, nil
}

func (us *medicineInteractor) Save(u *model.Medicine) (*model.Medicine, error) {
	data, err := us.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		u, err := us.MedicineRepository.Save(u)

		// do mailing
		// do logging
		// do another process
		return u, err
	})
	medicine, ok := data.(*model.Medicine)

	if !ok {
		return nil, errors.New("cast error")
	}

	if !errors.Is(err, nil) {
		return nil, err
	}

	return medicine, nil
}

func (us *medicineInteractor) Delete(id string) (u *model.Medicine, err error) {

	if u, err = us.MedicineRepository.Delete(id); err != nil {
		return nil, err
	}

	return u, nil
}
