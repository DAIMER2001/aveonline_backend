package registry

import (
	"github.com/DAIMER2001/demo-rest-api/interface/controller"
	interfacePresenter "github.com/DAIMER2001/demo-rest-api/interface/presenter"
	interfaceRepository "github.com/DAIMER2001/demo-rest-api/interface/repository"
	"github.com/DAIMER2001/demo-rest-api/usecase/interactor"
	"github.com/DAIMER2001/demo-rest-api/usecase/presenter"
	"github.com/DAIMER2001/demo-rest-api/usecase/repository"
)

func (r *registry) NewPromotionController() controller.PromotionController {
	return controller.NewPromotionController(r.NewPromotionInteractor())
}

func (r *registry) NewPromotionInteractor() interactor.PromotionInteractor {
	return interactor.NewPromotionInteractor(r.NewPromotionRepository(), r.NewPromotionPresenter(), interfaceRepository.NewDBRepository(r.db))
}

func (r *registry) NewPromotionRepository() repository.PromotionRepository {
	return interfaceRepository.NewPromotionRepository(r.db)
}

func (r *registry) NewPromotionPresenter() presenter.PromotionPresenter {
	return interfacePresenter.NewPromotionPresenter()
}
