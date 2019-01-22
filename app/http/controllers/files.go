package controllers

import (
	"fmt"
	"github.com/dotuancd/ezserve/app"
	"github.com/dotuancd/ezserve/app/http/errors"
	"github.com/dotuancd/ezserve/app/http/res"
	"github.com/dotuancd/ezserve/app/models"
	"github.com/dotuancd/ezserve/app/models/file"
	"github.com/dotuancd/ezserve/app/pagination"
	. "github.com/dotuancd/ezserve/app/supports"
	"github.com/dotuancd/ezserve/app/supports/str"
	"io/ioutil"

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

)

type FileHandler struct {
	App *app.App
}

func (h *FileHandler) Index() errors.HandlerFunc {
	return func(c *gin.Context) error {
		var files []models.File
		user := GetLoggedInUser(c)
		query := h.App.DB.Model(models.File{}).Where(models.File{UserID: user.ID})

		p := pagination.GetParamsContext(c).Paginate(query, &files)

		c.JSON(200, p)
		//c.JSON(200, gin.H{
		//	"ok": true,
		//	"files": p.Items(),
		//})

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

		_, err := c.MultipartForm()
		//isMultipartForm := err != nil
		//isJSON := c.GetHeader("Content-Type") == "application/json"

		fh, err := c.FormFile("file")

		if err == http.ErrMissingFile {
			c.PostForm("content")
		}

		if err != nil {
			return ErrFileUploadIsInvalid
		}

		dest := file.GetStoragePath(fh.Filename)

		if err = os.MkdirAll(path.Dir(dest), 0644); err != nil {
			return ErrCannotCreateFolder
		}

		err = c.SaveUploadedFile(fh, dest)

		if err != nil {
			fmt.Printf("An error occur while saving uploaded file %s", err.Error())

			return ErrCannotSaveFile
		}

		f := &models.File{
			ID: str.Rand(file.IdLength),
			Path: dest,
			ContentType: mime.TypeByExtension(path.Ext(fh.Filename)),
			UserID: GetLoggedInUser(c).ID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name: fh.Filename,
			Visibility: c.DefaultPostForm("visibility", file.VisibilityPublic),
			Secret: c.PostForm("secret"),
		}

		h.App.DB.Save(f)

		success := res.NewSuccess(c)

		success.Extra("file", f)

		success.Send()

		return nil
	}
}


func createFileFromContent(filename string, content string) (path string, err error){

	ioutil.WriteFile(filename, []byte(content), 0644)

	return path, nil
}
