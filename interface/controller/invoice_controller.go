package controller

import (
	"errors"
	"fmt"
	http "net/http"

	"github.com/DAIMER2001/demo-rest-api/domain/model"
	"github.com/DAIMER2001/demo-rest-api/usecase/interactor"
	"github.com/gin-gonic/gin"
)

type invoiceController struct {
	invoiceInteractor interactor.InvoiceInteractor
}

type InvoiceController interface {
	FindAllInvoices(c *gin.Context)
	FindInvoice(c *gin.Context)
	CreateInvoice(c *gin.Context)
	UpdateInvoice(c *gin.Context)
	DeleteInvoice(c *gin.Context)
	FindBetweenDatesInvoices(c *gin.Context)
}

func NewInvoiceController(us interactor.InvoiceInteractor) InvoiceController {
	return &invoiceController{us}
}

// FindAllInvoices	godoc
// @Summary      	List invoices
// @Description  	get invoices
// @Tags         	invoices
// @Accept       	json
// @Produce      	json
// @Success      	200  {array}   model.Invoice
// @Failure      	400  {string} string "error"
// @Failure      	404  {string} string "error"
// @Failure      	500  {string} string "error"
// @Router       	/invoice [get]
func (uc *invoiceController) FindAllInvoices(c *gin.Context) {
	fmt.Println("METHOD GET FindAllInvoices")
	var u []*model.Invoice

	u, err := uc.invoiceInteractor.Get(u)

	if err != nil {
		SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, u)
}

// FindBetweenDatesInvoices	godoc
// @Summary      			List invoices
// @Description  			get invoices for dates
// @Tags         			invoices
// @Accept       			json
// @Produce      			json
// @Success      			200  {array}   model.Invoice
// @Failure      			400  {string} string "error"
// @Failure      			404  {string} string "error"
// @Failure      			500  {string} string "error"
// @Router       			/invoice_dates [get]
func (uc *invoiceController) FindBetweenDatesInvoices(c *gin.Context) {
	fmt.Println("METHOD GET FindBetweenDatesInvoices")
	// var StartDate time.Time
	// var EndDate time.Time
	fmt.Println(c.Param("start_date"))
	fmt.Println(c.Param("end_date"))
	fmt.Println(c.Params.ByName("start_date"))
	fmt.Println(c.Params.ByName("end_date"))
	fmt.Println(c.Params)

	fmt.Println("not")
	var u []*model.Invoice
	u, err := uc.invoiceInteractor.FindForDates(c.Param("start_date"), c.Param("end_date"))

	if err != nil {
		SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, u)
}

// FindInvoice godoc
// @Summary      List invoices
// @Description  get invoices
// @Tags         invoices
// @Accept       json
// @Produce      json
// @Param        invoice_id   path      int  true  "invoice_id"
// @Success      200  {array}   model.Invoice
// @Failure      400  {string} string "error"
// @Failure      404  {string} string "error"
// @Failure      500  {string} string "error"
// @Router       /invoice/{id} [get]
func (uc *invoiceController) FindInvoice(c *gin.Context) {
	var u *model.Invoice

	if u, err := uc.invoiceInteractor.Find(u, c.Param("id")); err != nil {
		SendError(c, err)
		return
	} else {
		c.JSON(http.StatusOK, u)
	}

}

// CreateInvoice 	godoc
// @Summary 		Creates a invoice
// @Description 	creates Resource directory
// @Tags 			Invoices
// @Param 			fecha_crear path string true "date_payment"
// @Param 			pago_total path int true "full_payment"
// @Param 			medicines path array true "medicines"
// @Param 			promotions path array true "promotions"
// @Accept  		json
// @Success      	200  {object}  model.Invoice
// @Failure      	400  {string} string "error"
// @Failure      	404  {string} string "error"
// @Failure      	500  {string} string "error"
// @Router 			/invoice [post]
func (uc *invoiceController) CreateInvoice(c *gin.Context) {
	var params model.Invoice

	if err := c.Bind(&params); !errors.Is(err, nil) {
		SendError(c, err)
		return
	}

	if err := validateParams(c, &params); !errors.Is(err, nil) {
		SendError(c, err)
		return
	}

	if params, err := uc.invoiceInteractor.Save(&params); !errors.Is(err, nil) {
		SendError(c, err)
		return
	} else {
		c.JSON(http.StatusCreated, params)
	}

}

// CreateInvoiceMedicines 		godoc
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
func (uc *invoiceController) CreateInvoiceMedicines(c *gin.Context) {
	var params model.InvoiceMedicines

	if err := c.Bind(&params); !errors.Is(err, nil) {
		SendError(c, err)
		return
	}

	if err := validateParams(c, &params); !errors.Is(err, nil) {
		SendError(c, err)
		return
	}

	if params, err := uc.invoiceInteractor.SaveInvoiceMedicines(&params); !errors.Is(err, nil) {
		fmt.Println("err")
		fmt.Println(err.Error())
		SendError(c, err)
		return
	} else {
		c.JSON(http.StatusCreated, params)
	}
}

// UpdateInvoice  	godoc
// @Summary 		Updates a invoice
// @Description 	creates Resource directory
// @Tags 			Invoices
// @Param 			invoice_id path int true "invoice_id"
// @Param 			fecha_crear path string true "date_payment"
// @Param 			pago_total path int true "full_payment"
// @Param 			medicines path array true "medicines"
// @Param 			promotions path array true "promotions"
// @Accept  		json
// @Success      	200  {object}  model.Invoice
// @Failure      	400  {string} string "error"
// @Failure      	404  {string} string "error"
// @Failure      	500  {string} string "error"
// @Router 			/invoice [post]
func (uc *invoiceController) UpdateInvoice(c *gin.Context) {
	var paramsUpdate model.Invoice
	var paramsFound model.Invoice

	if err := c.Bind(&paramsUpdate); !errors.Is(err, nil) {
		SendError(c, err)
		return
	}

	if paramsFound, err := uc.invoiceInteractor.Find(&paramsFound, c.Param("id")); err != nil || paramsFound == nil {
		SendError(c, err)
		return
	}

	if err := validateParams(c, &paramsUpdate); !errors.Is(err, nil) {
		SendError(c, err)
		return
	}

	if paramsUpdate, err := uc.invoiceInteractor.Save(&paramsUpdate); !errors.Is(err, nil) {
		SendError(c, err)
		return
	} else {
		c.JSON(http.StatusCreated, paramsUpdate)
	}

}

// DeleteInvoice  godoc
// @Summary       	Delete a invoice
// @Description   	creates Resource directory
// @Tags 		  	Invoices
// @Accept  	  	json
// @Success       	200  {object}  model.Invoice
// @Failure       	400  {string} string "error"
// @Failure       	404  {string} string "error"
// @Failure       	500  {string} string "error"
// @Router 			/invoice{id} [delete]
func (uc *invoiceController) DeleteInvoice(c *gin.Context) {

	if u, err := uc.invoiceInteractor.Delete(c.Param("id")); !errors.Is(err, nil) {
		SendError(c, err)
		return
	} else {
		c.JSON(http.StatusOK, u)
	}
}
