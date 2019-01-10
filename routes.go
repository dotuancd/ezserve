package main

import (
	"github.com/dotuancd/ezserve/app"
	"github.com/dotuancd/ezserve/controllers"
	"github.com/dotuancd/ezserve/middlewares"
)

func registerFileRoutes(a *app.App) {

	files := &controllers.FileHandler{
		App: a,
	}

	authRoutes := a.Routes.Group("/", middlewares.UserAuth(a))
	authRoutes.GET("/api/files", files.Index())
	authRoutes.POST("/api/files", files.Store())

	a.Routes.GET("/files/:file_id/:filename", files.Show())
}