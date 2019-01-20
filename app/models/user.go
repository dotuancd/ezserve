package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string
	Email string `gorm:"not null;unique"`
	Password string
	ApiToken string `sql:"index"`
}