package controller

import (
	"errors"
	"fmt"
	http "net/http"

	"github.com/DAIMER2001/demo-rest-api/domain/model"
	"github.com/DAIMER2001/demo-rest-api/usecase/interactor"
	"github.com/gin-gonic/gin"
)

type promotionController struct {
	promotionInteractor interactor.PromotionInteractor
}

type PromotionController interface {
	FindAllPromotions(c *gin.Context)
	FindPromotion(c *gin.Context)
	CreatePromotion(c *gin.Context)
	UpdatePromotion(c *gin.Context)
	DeletePromotion(c *gin.Context)
	CreatePromotionMedicines(c *gin.Context)
}

func NewPromotionController(us interactor.PromotionInteractor) PromotionController {
	return &promotionController{us}
}

// FindAllPromotions 	godoc
// @Summary      		List promotions
// @Description  		get promotions
// @Tags         		promotions
// @Accept       		json
// @Produce      		json
// @Success      		200  {array}   model.Promotion
// @Failure      		400  {string} string "error"
// @Failure      		404  {string} string "error"
// @Failure      		500  {string} string "error"
// @Router       		/promotion [get]
func (uc *promotionController) FindAllPromotions(c *gin.Context) {
	var u []*model.Promotion

	u, err := uc.promotionInteractor.Get(u)

	if err != nil {
		SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, u)
}

// FindPromotion godoc
// @Summary      List promotions
// @Description  get promotions
// @Tags         promotions
// @Accept       json
// @Produce      json
// @Param        promotion_id   path      int  true  "promotion_id"
// @Success      200  {array}   model.Promotion
// @Failure      400  {string} string "error"
// @Failure      404  {string} string "error"
// @Failure      500  {string} string "error"
// @Router       /promotion/{id} [get]
func (uc *promotionController) FindPromotion(c *gin.Context) {

	var u *model.Promotion

	if u, err := uc.promotionInteractor.Find(u, c.Param("id")); err != nil {
		SendError(c, err)
		return
	} else {
		c.JSON(http.StatusOK, u)
	}

}

// CreatePromotion 	godoc
// @Summary 		Creates a promotion
// @Description 	creates Resource directory
// @Tags 			Promotions
// @Param 			description path string true "description"
// @Param 			percentage path int true "percentage"
// @Param 			start_date path string true "start_date"
// @Param 			end_date path string true "end_date"
// @Param 			promotions path array true "promotions"
// @Accept  		json
// @Success      	200  {object}  model.Promotion
// @Failure      	400  {string} string "error"
// @Failure      	404  {string} string "error"
// @Failure      	500  {string} string "error"
// @Router 			/promotion [post]
func (uc *promotionController) CreatePromotion(c *gin.Context) {
	var params model.Promotion

	if err := c.Bind(&params); !errors.Is(err, nil) {
		fmt.Println("err1")
		fmt.Println(err.Error())
		SendError(c, err)
		return
	}

	if err := validateParams(c, &params); !errors.Is(err, nil) {
		fmt.Println("err1")
		fmt.Println(err.Error())
		SendError(c, err)
		return
	}

	if params, err := uc.promotionInteractor.Save(&params); !errors.Is(err, nil) {
		fmt.Println("err")
		fmt.Println(err.Error())
		SendError(c, err)
		return
	} else {
		c.JSON(http.StatusCreated, params)
	}

}

// CreatePromotionMedicines 	godoc
// @Summary 					Creates a promotion
// @Description 				creates Resource directory
// @Tags 						Promotions
// @Param 						medicine_id path int true "medicine_id"
// @Param 						promotion_id path int true "promotion_id"
// @Accept  					json
// @Success      				200  {object}  model.Promotion
// @Failure      				400  {string} string "error"
// @Failure      				404  {string} string "error"
// @Failure      				500  {string} string "error"
// @Router 						/promotion_medicines [post]
func (uc *promotionController) CreatePromotionMedicines(c *gin.Context) {
	var params model.PromotionMedicines

	if err := c.Bind(&params); !errors.Is(err, nil) {
		SendError(c, err)
		return
	}

	if err := validateParams(c, &params); !errors.Is(err, nil) {
		SendError(c, err)
		return
	}

	if params, err := uc.promotionInteractor.SavePromotionMedicines(&params); !errors.Is(err, nil) {
		fmt.Println("err")
		fmt.Println(err.Error())
		SendError(c, err)
		return
	} else {
		c.JSON(http.StatusCreated, params)
	}

}

// UpdatePromotion  godoc
// @Summary 		Updates a promotion
// @Description 	creates Resource directory
// @Tags 			Promotions
// @Param 			promotion_id path int true "promotion_id"
// @Param 			description path string true "description"
// @Param 			percentage path int true "percentage"
// @Param 			start_date path string true "start_date"
// @Param 			end_date path string true "end_date"
// @Param 			promotions path array true "promotions"
// @Accept  		json
// @Success      	200  {object}  model.Promotion
// @Failure      	400  {string} string "error"
// @Failure      	404  {string} string "error"
// @Failure      	500  {string} string "error"
// @Router 			/promotion [post]
func (uc *promotionController) UpdatePromotion(c *gin.Context) {
	var paramsUpdate model.Promotion
	var paramsFound model.Promotion

	if err := c.Bind(&paramsUpdate); !errors.Is(err, nil) {
		SendError(c, err)
		return
	}

	if paramsFound, err := uc.promotionInteractor.Find(&paramsFound, c.Param("id")); err != nil || paramsFound == nil {
		SendError(c, err)
		return
	}

	if err := validateParams(c, &paramsUpdate); !errors.Is(err, nil) {
		SendError(c, err)
		return
	}

	if paramsUpdate, err := uc.promotionInteractor.Save(&paramsUpdate); !errors.Is(err, nil) {
		SendError(c, err)
		return
	} else {
		c.JSON(http.StatusCreated, paramsUpdate)
	}

}

// DeletePromotion  godoc
// @Summary       	Delete a promotion
// @Description   	creates Resource directory
// @Tags 		  	Promotions
// @Accept  	  	json
// @Success       	200  {object}  model.Promotion
// @Failure       	400  {string} string "error"
// @Failure       	404  {string} string "error"
// @Failure       	500  {string} string "error"
// @Router 			/promotion{id} [delete]
func (uc *promotionController) DeletePromotion(c *gin.Context) {

	if u, err := uc.promotionInteractor.Delete(c.Param("id")); !errors.Is(err, nil) {
		SendError(c, err)
		return
	} else {
		c.JSON(http.StatusOK, u)
	}
}
