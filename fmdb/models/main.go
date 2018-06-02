package models

import "github.com/jinzhu/gorm"

func Update(db *gorm.DB) (error) {
	if err := db.AutoMigrate(&Metadata{}).Error; err != nil {
		return err
	}
	if err := db.AutoMigrate(&Group{}).Error; err != nil {
		return err
	}
	if err := db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	return nil
}
