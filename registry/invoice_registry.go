package registry

import (
	"github.com/DAIMER2001/demo-rest-api/interface/controller"
	interfacePresenter "github.com/DAIMER2001/demo-rest-api/interface/presenter"
	interfaceRepository "github.com/DAIMER2001/demo-rest-api/interface/repository"
	"github.com/DAIMER2001/demo-rest-api/usecase/interactor"
	useCasePresenter "github.com/DAIMER2001/demo-rest-api/usecase/presenter"
	useCaseRepository "github.com/DAIMER2001/demo-rest-api/usecase/repository"
)

func (r *registry) NewInvoiceController() controller.InvoiceController {
	return controller.NewInvoiceController(r.NewInvoiceInteractor())
}

func (r *registry) NewInvoiceInteractor() interactor.InvoiceInteractor {
	return interactor.NewInvoiceInteractor(r.NewInvoiceRepository(), r.NewInvoicePresenter(), interfaceRepository.NewDBRepository(r.db))
}

func (r *registry) NewInvoiceRepository() useCaseRepository.InvoiceRepository {
	return interfaceRepository.NewInvoiceRepository(r.db)
}

func (r *registry) NewInvoicePresenter() useCasePresenter.InvoicePresenter {
	return interfacePresenter.NewInvoicePresenter()
}
