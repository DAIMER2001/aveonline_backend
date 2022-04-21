package registry

import (
	"github.com/DAIMER2001/demo-rest-api/interface/controller"
	interfacePresenter "github.com/DAIMER2001/demo-rest-api/interface/presenter"
	interfaceRepository "github.com/DAIMER2001/demo-rest-api/interface/repository"
	"github.com/DAIMER2001/demo-rest-api/usecase/interactor"
	"github.com/DAIMER2001/demo-rest-api/usecase/presenter"
	"github.com/DAIMER2001/demo-rest-api/usecase/repository"
)

func (r *registry) NewMedicineController() controller.MedicineController {
	return controller.NewMedicineController(r.NewMedicineInteractor())
}

func (r *registry) NewMedicineInteractor() interactor.MedicineInteractor {
	return interactor.NewMedicineInteractor(r.NewMedicineRepository(), r.NewMedicinePresenter(), interfaceRepository.NewDBRepository(r.db))
}

func (r *registry) NewMedicineRepository() repository.MedicineRepository {
	return interfaceRepository.NewMedicineRepository(r.db)
}

func (r *registry) NewMedicinePresenter() presenter.MedicinePresenter {
	return interfacePresenter.NewMedicinePresenter()
}
