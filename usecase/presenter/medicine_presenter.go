package presenter

import "github.com/DAIMER2001/demo-rest-api/domain/model"

type MedicinePresenter interface {
	ResponseMedicines(u []*model.Medicine) []*model.Medicine
}
