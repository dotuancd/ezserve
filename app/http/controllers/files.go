package controllers

import (
	"github.com/dotuancd/ezserve/app"
	"github.com/dotuancd/ezserve/app/http/errors"
	"github.com/dotuancd/ezserve/app/http/res"
	"github.com/dotuancd/ezserve/app/models"
	"github.com/dotuancd/ezserve/app/models/file"
	"github.com/dotuancd/ezserve/app/pagination"
	. "github.com/dotuancd/ezserve/app/supports"
	"github.com/dotuancd/ezserve/app/supports/str"
	"io"
	"mime/multipart"
	"strings"

	"github.com/gin-gonic/gin"

	"mime"
	"net/http"
	"os"
	"path"
	"time"
)

var (
	ErrFileNotFound = errors.New("file_not_found", "File not found", http.StatusNotFound)

	ErrFileCannotRead = errors.New("cannot_read_file", "An error occur while reading file", http.StatusInternalServerError)

	ErrFileUploadIsInvalid = errors.New("cannot_read_uploaded_file", "An error occur while reading the upload file", http.StatusUnprocessableEntity)

	ErrCannotCreateFolder = errors.New("cannot_make_dir", "An error occur while creating directories for uploaded file", http.StatusInternalServerError)

	ErrCannotSaveFile = errors.New("cannot_save_uploaded_file", "An error occur while saving uploaded file", http.StatusInternalServerError)

	ErrContentIsRequired = errors.New("cannot_save_uploaded_file", "An error occur while saving uploaded file", http.StatusInternalServerError)

)

type FileHandler struct {
	App *app.App
}

type StoreRequest struct {
	File         *multipart.FileHeader
	Filename     string `json:"filename" form:"filename"`
	Content      string `json:"content" form:"content"`
	Visibility   string `json:"visibility" form:"visibility"`
	Secret       string `json:"secret" form:"secret"`

	isFileUpload bool
}

func (h *FileHandler) SomeAction(c *gin.Context) error  {

	return nil
}

func (h *FileHandler) Index() errors.HandlerFunc {
	return func(c *gin.Context) error {
		var files []models.File
		user := GetLoggedInUser(c)
		query := h.App.DB.Model(models.File{}).Where(models.File{UserID: user.ID})

		p := pagination.GetParamsContext(c).Paginate(query, &files)

		c.JSON(200, p)

		return nil
	}
}

func (h *FileHandler) Show() errors.HandlerFunc {
	return func(c *gin.Context) error {
		id := c.Param("file_id")
		f := models.File{}
		h.App.DB.Find(&f, &models.File{ID: id})

		if f.ID == "" {
			return ErrFileNotFound
		}

		_, err := os.Open(f.Path)

		if err != nil {
			return ErrFileCannotRead
		}

		c.Writer.Header().Set("Content-Type", f.ContentType)
		c.Writer.Header().Set("Content-Disposition", "inline")

		c.File(f.Path)

		return nil
	}
}

func (h *FileHandler) Store () errors.HandlerFunc {
	return func(c *gin.Context) error {
		// Validate, binding request
		rq, err := getSaveRequest(c)

		if err != nil {
			return err
		}

		// Get destination filepath
		dest := file.GetStoragePath(rq.Filename)

		// Ensure parent folder are exists
		if err = os.MkdirAll(path.Dir(dest), 0644); err != nil {
			return ErrCannotCreateFolder
		}

		if rq.isFileUpload {
			err = c.SaveUploadedFile(rq.File, dest)
		} else {
			err = saveContentAsFile(rq.Content, dest)
		}

		if err != nil {
			return ErrCannotSaveFile
		}

		contentType := mime.TypeByExtension(path.Ext(rq.Filename))

		f := &models.File{
			ID: str.Rand(file.IdLength),
			Path: dest,
			ContentType: contentType,
			UserID: GetLoggedInUser(c).ID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name: rq.Filename,
			Visibility: rq.Visibility,
			Secret: rq.Secret,
		}

		h.App.DB.Save(f)

		success := res.NewSuccess(c)

		success.Extra("file", f)

		success.Send()

		return nil
	}
}

func getSaveRequest(c *gin.Context) (*StoreRequest, error) {
	rq := &StoreRequest{}
	err := c.ShouldBind(rq)

	if err != nil {
		return nil, err
	}

	rq.File, err = c.FormFile("file")
	rq.isFileUpload = rq.File != nil

	// validate
	if !rq.isFileUpload && rq.Content == "" {
		return nil, requiredWhenFileNotRepresent("content")
	}

	if !rq.isFileUpload && rq.Filename == "" {
		return nil, requiredWhenFileNotRepresent("filename")
	}

	// default value
	if rq.isFileUpload && rq.Filename == "" {
		rq.Filename = rq.File.Filename
	}

	if rq.Visibility == "" {
		rq.Visibility = file.VisibilityPublic
	}

	return rq, nil
}

func saveContentAsFile(content string, dest string) error {

	w, err := os.Create(dest)

	if err != nil {
		return err
	}

	defer Close(w)

	r := strings.NewReader(content)

	_, err = io.Copy(w, r)

	return err
}

func requiredWhenFileNotRepresent(fieldName string) *errors.Error {
	err := errors.New("validation_failed", "Validation failed", http.StatusUnprocessableEntity)
	err.Fields = map[string]interface{}{
		fieldName: str.Replacements("The :attribute is required when file not represent", map[string]interface{}{
			"attribute": fieldName,
		}),
	}

	return err
}