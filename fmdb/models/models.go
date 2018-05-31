package models

type Metadata struct {
	Key string `gorm:"type:varchar(255);PRIMARY_KEY"`
	Value string
}

type UserSource int

const (
	UserSourceLocal = 0
	UserSourceLDAP = 1
)



type User struct {
	Username string `gorm:"type:varchar(255);PRIMARY_KEY"`
	Source   UserSource
	IsAdmin  bool
	IsActive bool

}