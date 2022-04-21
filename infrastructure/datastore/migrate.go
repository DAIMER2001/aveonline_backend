package datastore

import (
	"log"

	"github.com/DAIMER2001/demo-rest-api/domain/model"
	"gorm.io/gorm"
)

func Migrations(db *gorm.DB) error {

	log.Println("MIGRANDO....")

	db.AutoMigrate(&model.Invoice{})
	db.AutoMigrate(&model.Medicine{})
	db.AutoMigrate(&model.Promotion{})
	return nil
}
