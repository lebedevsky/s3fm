package fmdb

import (
	"github.com/lebedevsky/s3fm/fmdb/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
)

func OpenDB(path string) (db *gorm.DB, err error) {
	db, err = gorm.Open("sqlite3", path)
	if err != nil {
		return
	}

	if err = models.Update(db); err != nil {
		return
	}

	status := &models.Metadata{}
	db.Where("key = ?", "initialized").First(status)
	if status.Value == "" {
		if err = initDB(db); err != nil {
			return
		}
		db.Create(&models.Metadata{Key: "initialized", Value: "true"})
	}

	return
}

func initDB(db *gorm.DB) error {
	fmt.Printf("Time to init DB\n")
	db.Create(&models.Metadata{Key: "db_version", Value: "0"})
	db.Create(&models.User{Username: "admin", Source: models.UserSourceLocal, IsAdmin: true, IsActive: true})
	return nil
}