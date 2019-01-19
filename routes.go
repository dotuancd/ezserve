package main

import (
	"github.com/dotuancd/ezserve/app"
	"github.com/dotuancd/ezserve/controllers"
	"github.com/dotuancd/ezserve/http/errors"
	"github.com/dotuancd/ezserve/middlewares"
)

func registerFileRoutes(a *app.App) {

	home := controllers.HomeController{
		App:a,
	}

	a.Routes.GET("/", errors.Handler(home.Index))

	files := &controllers.FileHandler{
		App: a,
	}

	authRoutes := a.Routes.Group("/", middlewares.UserAuth(a))
	authRoutes.GET("/api/files", errors.Handler(files.Index()))
	authRoutes.POST("/api/files", errors.Handler(files.Store()))

	a.Routes.GET("/files/:file_id/:filename", errors.Handler(files.Show()))
}