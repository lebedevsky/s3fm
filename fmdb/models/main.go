package models

import "github.com/jinzhu/gorm"

func Update(db *gorm.DB) (error) {
	db.AutoMigrate(&Metadata{})
	db.AutoMigrate(&User{})
	return nil
}
