package controller

import (
	"errors"
	"fmt"
	http "net/http"

	"github.com/DAIMER2001/demo-rest-api/domain/model"
	"github.com/DAIMER2001/demo-rest-api/usecase/interactor"
	"github.com/gin-gonic/gin"
)

type medicineController struct {
	medicineInteractor interactor.MedicineInteractor
}

type MedicineController interface {
	FindAllMedicines(c *gin.Context)
	FindMedicine(c *gin.Context)
	CreateMedicine(c *gin.Context)
	UpdateMedicine(c *gin.Context)
	DeleteMedicine(c *gin.Context)
}

func NewMedicineController(us interactor.MedicineInteractor) MedicineController {
	return &medicineController{us}
}

// GetMedicines godoc
// @Summary      List medicines
// @Description  get medicines
// @Tags         medicines
// @Accept       json
// @Produce      json

// @Success      200  {array}   model.Medicine
// @Failure      400  {string} string "error"
// @Failure      404  {string} string "error"
// @Failure      500  {string} string "error"
// @Router       /medicine [get]
func (uc *medicineController) FindAllMedicines(c *gin.Context) {
	var u []*model.Medicine

	fmt.Println(&u)

	u, err := uc.medicineInteractor.Get(u)
	fmt.Println(u)
	fmt.Print(u)
	fmt.Print(u)

	if err != nil {
		SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, u)
}

// FindMedicine godoc
// @Summary      List medicines
// @Description  get medicines
// @Tags         medicines
// @Accept       json
// @Produce      json
// @Param        medicine_id   path      int  true  "medicine_id"
// @Success      200  {array}   model.Medicine
// @Failure      400  {string} string "error"
// @Failure      404  {string} string "error"
// @Failure      500  {string} string "error"
// @Router       /medicine/{id} [get]
func (uc *medicineController) FindMedicine(c *gin.Context) {

	var u *model.Medicine

	if u, err := uc.medicineInteractor.Find(u, c.Param("id")); err != nil {
		SendError(c, err)
		return
	} else {
		c.JSON(http.StatusOK, u)
	}

}

// CreateMedicine godoc
// @Summary Creates a medicine
// @Description creates Resource directory
// @Tags Medicines
// @Param name path string true "name"
// @Param price path int true "price"
// @Param location path string true "location"
// @Accept  json
// @Success      200  {object}  model.Medicine
// @Failure      400  {string} string "error"
// @Failure      404  {string} string "error"
// @Failure      500  {string} string "error"
// @Router /medicine [post]
func (uc *medicineController) CreateMedicine(c *gin.Context) {
	var params model.Medicine

	if err := c.Bind(&params); !errors.Is(err, nil) {
		SendError(c, err)
		return
	}

	if err := validateParams(c, &params); !errors.Is(err, nil) {
		SendError(c, err)
		return
	}

	if params, err := uc.medicineInteractor.Save(&params); !errors.Is(err, nil) {
		SendError(c, err, "error al guardar los datos")
		return
	} else {
		c.JSON(http.StatusCreated, params)
	}

}

// UpdateMedicine  godoc
// @Summary 		Updates a medicine
// @Description 	creates Resource directory
// @Tags 			Medicines
// @Param        	medicine_id   path      int  true  "medicine_id"
// @Param 			name path string true "name"
// @Param 			price path int true "price"
// @Param 			location path string true "location"
// @Accept  		json
// @Success      	200  {object}  model.Medicine
// @Failure      	400  {string} string "error"
// @Failure      	404  {string} string "error"
// @Failure      	500  {string} string "error"
// @Router 			/medicine [post]
func (uc *medicineController) UpdateMedicine(c *gin.Context) {
	var paramsUpdate model.Medicine
	var paramsFound model.Medicine

	if err := c.Bind(&paramsUpdate); !errors.Is(err, nil) {
		SendError(c, err)
		return
	}

	if paramsFound, err := uc.medicineInteractor.Find(&paramsFound, c.Param("id")); err != nil || paramsFound == nil {
		SendError(c, err)
		return
	}

	if err := validateParams(c, &paramsUpdate); !errors.Is(err, nil) {
		SendError(c, err)
		return
	}

	if paramsUpdate, err := uc.medicineInteractor.Save(&paramsUpdate); !errors.Is(err, nil) {
		SendError(c, err)
		return
	} else {
		c.JSON(http.StatusCreated, paramsUpdate)
	}

}

// DeleteMedicine  godoc
// @Summary       	Delete a medicine
// @Description   	creates Resource directory
// @Tags 		  	Medicines
// @Accept  	  	json
// @Success       	200  {object}  model.Medicine
// @Failure       	400  {string} string "error"
// @Failure       	404  {string} string "error"
// @Failure       	500  {string} string "error"
// @Router 			/medicine{id} [delete]
func (uc *medicineController) DeleteMedicine(c *gin.Context) {

	if u, err := uc.medicineInteractor.Delete(c.Param("id")); !errors.Is(err, nil) {
		SendError(c, err)
		return
	} else {
		c.JSON(http.StatusOK, u)
	}
}
