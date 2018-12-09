package controllers

import (
	"bitbucket.org/ezserve/ezserve/app"
	"bitbucket.org/ezserve/ezserve/http/res"
	"bitbucket.org/ezserve/ezserve/models"
	"bitbucket.org/ezserve/ezserve/supports"
	"crypto/sha1"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime"
	"os"
	"path"
	"time"
)

type FileHandler struct {
	App *app.App
}

var root = "storage"

func (h *FileHandler) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		var files []models.File
		user := c.MustGet("user").(models.User)

		db := h.App.DB

		db.Find(&files, models.File{UserID: user.ID})

		c.JSON(200, gin.H{
			"ok": true,
			"files": files,
		})
	}
}

func (h *FileHandler) Show() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("file_id")
		file := models.File{}
		h.App.DB.Find(&file, &models.File{ID: id})

		if file.ID == "" {
			fileNotFound(c)
			return
		}

		c.Writer.Header().Set("Content-Type", file.ContentType)
		c.File(file.Path)
	}
}

func (h *FileHandler) Store () gin.HandlerFunc {
	return func(c *gin.Context) {
		fh, err := c.FormFile("file")

		if err != nil {
			res.SendError(
				c,
				res.HttpBadRequest,
				"invalid_file_upload",
				"Cannot read the uploaded file",
			)
			return
		}

		dest, _ := getStoragePath(fh.Filename)

		if err = os.MkdirAll(path.Dir(dest), 0644); err != nil {
			res.SendError(
				c,
				res.HttpInternalError,
				"cannot_make_dir",
				"An error occur while creating directories for uploaded file",
			)
			return
		}

		err = c.SaveUploadedFile(fh, dest)

		if err != nil {
			fmt.Printf("An error occur while saving uploaded file %s", err.Error())
			res.SendError(
				c,
				res.HttpInternalError,
				"cannot_save_uploaded_file",
				"An error occur while saving uploaded file",
			)
			return
		}

		file := &models.File{
			ID: supports.StringRand(models.FileIdLength),
			Path: dest,
			ContentType: mime.TypeByExtension(path.Ext(fh.Filename)),
			UserID: toUser(c).ID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		h.App.DB.Save(file)

		success := res.NewSuccess(c)

		success.Content()["file"] = file

		success.Send()

	}
}

func getStoragePath(filename string) (string, error) {
	hash := sha1.New()
	hash.Write([]byte(time.Now().String() + filename))
	s := fmt.Sprintf("%x", hash.Sum(nil))

	return path.Join(root, s[:3], s[3:6], s[6:] + "-" + filename), nil
}

//func moveUploadedFile(c *gin.Context, name string, dest string) (error) {
//
//	uploadedFile, err := c.FormFile(name)
//
//	if err != nil {
//		return err
//	}
//
//	if err := c.SaveUploadedFile(uploadedFile, dest); err != nil {
//		return err
//	}
//
//	return nil
//}

func toUser(c *gin.Context) models.User {
	return c.MustGet("user").(models.User)
}

func fileNotFound(c *gin.Context) {
	res.
		NotFound(c).
		Code("file_not_found").
		Message("File not found").
		Send()
	return
}
