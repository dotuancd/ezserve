package models

import (
	"bitbucket.org/ezserve/ezserve/supports"
	"github.com/spf13/viper"
	"time"
)

const FileIdLength = 12

const (
	FileVisibilityPublic = "public"
	FileVisibilityProtected = "protected"
	FileVisibilityPrivate = "private"
)

type File struct {
	ID string `json:"id" gorm:"primary_key;size:30"`
	ContentType string `json:"content_type"`
	Path string `json:"path"`
	Name string `json:"name"`
	Secret string `json:"secret"`
	UserID uint `json:"user_id" sql:"index"`
	Visibility string `json:"visibility" gorm:"default:public;size:20"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
	PublicURL string `json:"public_url" gorm:"-"`
}

func (file *File) AfterFind() error {
	return file.appendPublicURL()
}

func (file *File) AfterSave() error {
	return file.appendPublicURL()
}

func (file *File) appendPublicURL() error {
	replacements := map[string]interface{}{
		"host": viper.GetString("app.host"),
		"file_id": file.ID,
		"filename": file.Name,
	}

	file.PublicURL = supports.Replace(
		":host/files/:file_id/:filename",
		replacements,
	)

	return nil
}
