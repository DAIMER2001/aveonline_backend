package router

import (
	"fmt"
	"log"
	"os"

	"github.com/DAIMER2001/demo-rest-api/config"
	"github.com/DAIMER2001/demo-rest-api/interface/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter(
	server *gin.Engine,
	appController controller.AppController,
) *gin.Engine {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("AVEONLINE_PORT")

	_ = server.SetTrustedProxies([]string{})

	server.Use(cors.Default())
	server.GET("/", func(c *gin.Context) {
		c.String(200, " WELCOME AVEONLINE")
	})

	v1 := server.Group("/v1")
	{
		// INVOICE ROUTES
		v1.GET("/invoice", appController.Invoice.FindAllInvoices)
		v1.GET("/invoice_dates/:start_date/*end_date", appController.Invoice.FindBetweenDatesInvoices)
		v1.GET("/invoice/:id", appController.Invoice.FindInvoice)
		v1.POST("/invoice", appController.Invoice.CreateInvoice)
		v1.PUT("/invoice", appController.Invoice.UpdateInvoice)
		v1.DELETE("/invoice/:id", appController.Invoice.DeleteInvoice)

		// MEDICINE ROUTES
		v1.GET("/medicine", appController.Medicine.FindAllMedicines)
		v1.GET("/medicine/:id", appController.Medicine.FindMedicine)
		v1.POST("/medicine", appController.Medicine.CreateMedicine)
		v1.PUT("/medicine", appController.Medicine.UpdateMedicine)
		v1.DELETE("/medicine/:id", appController.Medicine.DeleteMedicine)

		// PROMOTION ROUTES
		v1.GET("/promotion", appController.Promotion.FindAllPromotions)
		v1.GET("/promotion/:id", appController.Promotion.FindPromotion)
		v1.POST("/promotion", appController.Promotion.CreatePromotion)
		v1.POST("/promotion_medicines", appController.Promotion.CreatePromotionMedicines)
		v1.PUT("/promotion", appController.Promotion.UpdatePromotion)
		v1.DELETE("/promotion/:id", appController.Promotion.DeletePromotion)
	}
	fmt.Println(port)

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	return server
}
