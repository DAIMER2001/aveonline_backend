package main

import (
	"fmt"
	"log"
	"time"

	"github.com/DAIMER2001/demo-rest-api/config"
	_ "github.com/DAIMER2001/demo-rest-api/docs"
	"github.com/DAIMER2001/demo-rest-api/infrastructure/datastore"
	"github.com/DAIMER2001/demo-rest-api/infrastructure/router"
	"github.com/DAIMER2001/demo-rest-api/registry"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8082
// @BasePath  /v1

// @securityDefinitions.basic  BasicAuth

type Server struct {
	Gin *gin.Engine
}

func main() {
	config.ReadConfig()

	port := config.C.Server.Address
	db := datastore.NewDB()

	r := registry.NewRegistry(db)

	err := datastore.Migrations(db)

	if err != nil {
		log.Fatal(err)
	}

	server := Server{}
	server.Gin = gin.New()

	server.Gin.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE, PATCH",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	server.Gin = router.NewRouter(
		server.Gin,
		r.NewAppController())

	// e := echo.New()
	// e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + port)
	if err := server.Gin.Run(":" + port); err != nil {
		log.Fatalln(err)
	}

	// commons.Migrate()

}
