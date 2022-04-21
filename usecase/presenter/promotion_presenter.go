package presenter

import "github.com/DAIMER2001/demo-rest-api/domain/model"

type PromotionPresenter interface {
	ResponsePromotions(u []*model.Promotion) []*model.Promotion
}
