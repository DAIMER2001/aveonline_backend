package presenter

import (
	"github.com/DAIMER2001/demo-rest-api/domain/model"
	"github.com/DAIMER2001/demo-rest-api/usecase/presenter"
)

type medicinePresenter struct{}

func NewMedicinePresenter() presenter.MedicinePresenter {
	return &medicinePresenter{}
}

func (up *medicinePresenter) ResponseMedicines(us []*model.Medicine) []*model.Medicine {
	return us
}
