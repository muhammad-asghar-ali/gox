package helpers

import (
	"sync"

	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
	oc sync.Once
)

func ConnectDB() *gorm.DB {
	oc.Do(func() {
		var err error
		connection := ConnectionString()
		db, err = gorm.Open("postgres", connection)
		HandleError(err)
	})

	return db
}

func GetDatabase() *gorm.DB {
	if db == nil {
		return ConnectDB()
	}

	return db
}
