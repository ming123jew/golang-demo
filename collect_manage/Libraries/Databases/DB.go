package Databases

import (
	"gorm.io/gorm"
)

type DB struct {
}

func NewDB(dialector gorm.Dialector, opts ...gorm.Option) (*gorm.DB, error) {

	db, err := gorm.Open(dialector, opts...)

	if err != nil {
		panic("failed to connect database")
	}

	return db, err
}
