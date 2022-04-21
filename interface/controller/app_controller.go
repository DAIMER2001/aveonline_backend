package controller

type AppController struct {
	Invoice   interface{ InvoiceController }
	Medicine  interface{ MedicineController }
	Promotion interface{ PromotionController }
}
