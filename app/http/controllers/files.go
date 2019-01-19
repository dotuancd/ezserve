package controllers

import (
	"crypto/sha1"
	"fmt"
	"github.com/dotuancd/ezserve/app/http/errors"

	"github.com/dotuancd/ezserve/app"
	"github.com/dotuancd/ezserve/app/http/res"
	"github.com/dotuancd/ezserve/app/models"
	"github.com/dotuancd/ezserve/app/supports"

	"github.com/gin-gonic/gin"

	"mime"
	"net/http"
	"os"
	"path"
	"time"
)

var (
	ErrFileNotFound = errors.New("file_not_found", "File not found", http.StatusNotFound)

	ErrFileUploadIsInvalid = errors.New("cannot_read_uploaded_file", "An error occur while reading the upload file", http.StatusUnprocessableEntity)

	ErrCannotCreateFolder = errors.New("cannot_make_dir", "An error occur while creating directories for uploaded file", http.StatusInternalServerError)

	ErrCannotSaveFile = errors.New("cannot_save_uploaded_file", "An error occur while saving uploaded file", http.StatusInternalServerError)

)

type FileHandler struct {
	App *app.App
}

var root = "storage"


func (h *FileHandler) Index() errors.HandlerFunc {
	return func(c *gin.Context) error {
		var files []models.File
		user := c.MustGet("user").(models.User)

		db := h.App.DB

		db.Find(&files, models.File{UserID: user.ID})

		c.JSON(200, gin.H{
			"ok": true,
			"files": files,
		})

		return nil
	}
}

func (h *FileHandler) Show() errors.HandlerFunc {
	return func(c *gin.Context) error {
		id := c.Param("file_id")
		file := models.File{}
		h.App.DB.Find(&file, &models.File{ID: id})

		if file.ID == "" {
			return ErrFileNotFound
		}

		c.Writer.Header().Set("Content-Type", file.ContentType)
		c.File(file.Path)

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
			c.GetHeader("content")
		}

		if err != nil {
			return ErrFileUploadIsInvalid
		}

		dest := getStoragePath(fh.Filename)

		if err = os.MkdirAll(path.Dir(dest), 0644); err != nil {
			return ErrCannotCreateFolder
		}

		err = c.SaveUploadedFile(fh, dest)

		if err != nil {
			fmt.Printf("An error occur while saving uploaded file %s", err.Error())

			return ErrCannotSaveFile
		}

		file := &models.File{
			ID: supports.StringRand(models.FileIdLength),
			Path: dest,
			ContentType: mime.TypeByExtension(path.Ext(fh.Filename)),
			UserID: toUser(c).ID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name: fh.Filename,
			Visibility: c.DefaultPostForm("visibility", models.FileVisibilityPrivate),
			Secret: c.PostForm("secret"),
		}

		h.App.DB.Save(file)

		success := res.NewSuccess(c)

		success.Extra("file", file)

		success.Send()

		return nil
	}
}

func getStoragePath(filename string) string {
	hash := sha1.New()
	hash.Write([]byte(time.Now().String() + filename))
	hashed := fmt.Sprintf("%x", hash.Sum(nil))

	return path.Join(root, hashed[:3], hashed[3:6], hashed[6:] + "-" + filename)
}

func toUser(c *gin.Context) models.User {
	return c.MustGet("user").(models.User)
}