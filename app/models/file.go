package models

import (
	"github.com/dotuancd/ezserve/app/supports/str"
	"github.com/spf13/viper"
	"time"
)

type File struct {
	ID string `json:"id" gorm:"primary_key;size:30"`
	ContentType string `json:"content_type"`
	Path string `json:"-"`
	Name string `json:"name"`
	Secret string `json:"secret"`
	UserID uint `json:"user_id" sql:"index"`
	Visibility string `json:"visibility" gorm:"default:'public';size:20"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
	PublicURL string `json:"public_url" gorm:"-"`
}

func (f *File) AfterFind() error {
	return f.appendPublicURL()
}

func (f *File) AfterSave() error {
	return f.appendPublicURL()
}

func (f *File) appendPublicURL() error {
	replacements := map[string]interface{}{
		"host":     viper.GetString("app.host"),
		"file_id":  f.ID,
		"filename": f.Name,
	}

	f.PublicURL = str.Replacements(
		":host/files/:file_id/:filename",
		replacements,
	)

	return nil
}
