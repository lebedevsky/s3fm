package fmdb

import (
	"github.com/lebedevsky/s3fm/fmdb/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
	"github.com/lebedevsky/s3fm/utils"
)

func OpenDB(path string) (db *gorm.DB, err error) {
	db, err = gorm.Open("sqlite3", path)
	if err != nil {
		return
	}

	if err = models.Update(db); err != nil {
		return
	}

	if err = db.Model(&models.Metadata{Key: "initialized"}).First(&models.Metadata{}).Error; err == gorm.ErrRecordNotFound {
		if err = initDB(db); err != nil {
			return
		}
		if err = db.Create(&models.Metadata{Key: "initialized", Value: "true"}).Error; err != nil {
			return
		}
	} else if err != nil {
		return
	}

	if err= checkAdminUser(db); err != nil {
		return
	}

	return
}

func initDB(db *gorm.DB) error {
	fmt.Printf("Time to init DB\n")
	db.Create(&models.Metadata{Key: "db_version", Value: "0"})
	db.Create(&models.Group{Name: "admin", Source: models.UserSourceLocal, IsAdmin: true})
	return nil
}

func checkAdminUser(db *gorm.DB) error {
	const (
		adminUserName  = "admin"
		adminGroupName = "admin"
	)
	user := models.User{Login: adminUserName}
	err := db.Where(&user).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	} else if err == gorm.ErrRecordNotFound  {
		user.Source = models.UserSourceLocal
		user.IsActive = true

		salt, err := utils.GetSalt()
		if err != nil {
			return err
		}
		user.Salt = salt

		pass, err := utils.GetRandomPassword()
		if err != nil {
			return err
		}
		passHash, err := utils.GetPasswordHash(pass, user.Salt)
		if err != nil {
			return err
		}
		user.PasswordHash = passHash

		group := models.Group{Name: adminGroupName}
		err = db.Where(&group).First(&group).Error
		if err != nil {
			return err
		}
		user.Groups = append(user.Groups, group)
		user.Groups = []models.Group{group,}
		if err = db.Create(&user).Error; err != nil {
			return err
		}

		fmt.Printf("User %s created, new password: %s\n", adminUserName, pass)
	}

	return nil
}