package controllers

import (
	"github.com/dotuancd/ezserve/app"
	"github.com/gin-gonic/gin"
)

type HomeController struct {
	App *app.App
}

func (h *HomeController) Index(c *gin.Context) error {

	return nil
}

//func (h *HomeController) Index () gin.HandlerFunc {
//	return HandleErrorFunc(func(c *gin.Context) error {
//
//		return nil
//	})
//}