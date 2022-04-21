package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/DAIMER2001/demo-rest-api/usecase/repository"
	"gorm.io/gorm"
)

type dbRepository struct {
	db *gorm.DB
}

func NewDBRepository(db *gorm.DB) repository.DBRepository {
	return &dbRepository{db}
}

func (r *dbRepository) Transaction(txFunc func(interface{}) (interface{}, error)) (data interface{}, err error) {
	tx := r.db.Begin()
	fmt.Println("TRANS")
	fmt.Println(tx)
	if !errors.Is(tx.Error, nil) {
		return nil, tx.Error
	}

	defer func() {
		if p := recover(); p != nil {
			log.Print("recover")
			tx.Rollback()
			panic(p)
		} else if !errors.Is(err, nil) {
			log.Print("rollback")
			tx.Rollback()
			panic("error")
		} else {
			err = tx.Commit().Error
			fmt.Println(err)
		}
	}()

	data, err = txFunc(tx)
	fmt.Println(data)
	fmt.Println(err)
	return data, err
}
