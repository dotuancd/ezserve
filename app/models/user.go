package models

import "time"

type User struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Name string `json:"name"`
	Email string `gorm:"not null;unique" json:"email"`
	Password string `json:"-"`
	ApiToken string `sql:"index" json:"api_token"`
}