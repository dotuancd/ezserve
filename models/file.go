package models

import (
	//"path"
	"time"
)

const FileIdLength = 30

type File struct {
	ID string `gorm:"primary_key;size:30";json:"id"`
	ContentType string `json:"content_type"`
	//Content []byte `gorm:"type:mediumblob";json:"content"`
	Path string `json:"path";gorm:"size:255"`
	Secret string `json:"secret"`
	UserID uint `sql:"index";json:"user_id";gorm:"column:user_id"`
	IsPrivate bool `json:"is_private"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index";json:"deleted_at"`
}

func (file *File) GetPath() string {
	//return path.Join("storage", file.Path)
	return ""
}
