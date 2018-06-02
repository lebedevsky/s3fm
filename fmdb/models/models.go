package models

import "github.com/jinzhu/gorm"

type UserSource int

const (
	UserSourceLocal UserSource = 0
	UserSourceLDAP UserSource = 1
)

type Metadata struct {
	Key string `gorm:"type:varchar(255);primary_key"`
	Value string
}

type User struct {
	Login        string `gorm:"type:varchar(255);primary_key"`
	Source       UserSource
	IsActive     bool
	PasswordHash string
	Salt         string
	Groups       []Group `gorm:"many2many:user_groups;"`
}

type Group struct {
	gorm.Model
	Name    string `gorm:"type:varchar(255);UNIQUE_INDEX"`
	Source  UserSource
	IsAdmin bool
}


