package datastore

import (
	"fmt"
	"log"

	"github.com/DAIMER2001/demo-rest-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {

	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable TimeZone=America/Bogota",
		config.C.Database.Host, config.C.Database.Port,
		config.C.Database.User, config.C.Database.Password, config.C.Database.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
