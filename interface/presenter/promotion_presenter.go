package presenter

import (
	"github.com/DAIMER2001/demo-rest-api/domain/model"
	"github.com/DAIMER2001/demo-rest-api/usecase/presenter"
)

type promotionPresenter struct{}

func NewPromotionPresenter() presenter.PromotionPresenter {
	return &promotionPresenter{}
}

func (up *promotionPresenter) ResponsePromotions(us []*model.Promotion) []*model.Promotion {
	return us
}
