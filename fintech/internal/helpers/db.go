package helpers

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(200)

	return db
}

func GetDatabase() *gorm.DB {
	if db == nil {
		return ConnectDB()
	}

	return db
}
