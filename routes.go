package main

import (
	"github.com/dotuancd/ezserve/app"
	"github.com/dotuancd/ezserve/controllers"
	"github.com/dotuancd/ezserve/http/errors"
	"github.com/dotuancd/ezserve/middlewares"
	"github.com/gin-gonic/gin"
)

func handlingErrorFunc(routes gin.IRoutes) {

}

func registerFileRoutes(a *app.App) {

	home := controllers.HomeController{
		App:a,
	}

	adater := errors.HandlerFuncAdapter{Next: home.Index()}

	a.Routes.AppEngine

	a.Routes.GET("/", adater.Handle())

	files := &controllers.FileHandler{
		App: a,
	}

	authRoutes := a.Routes.Group("/", middlewares.UserAuth(a))
	authRoutes.GET("/api/files", files.Index())
	authRoutes.POST("/api/files", files.Store())

	a.Routes.GET("/files/:file_id/:filename", files.Show())
}