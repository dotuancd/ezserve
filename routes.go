package main

import (
	"bitbucket.org/ezserve/ezserve/app"
	"bitbucket.org/ezserve/ezserve/controllers"
	"bitbucket.org/ezserve/ezserve/middlewares"
)

func registerFileRoutes(a *app.App) {

	files := &controllers.FileHandler{
		App: a,
	}

	authRoutes := a.Routes.Group("/", middlewares.UserAuth(a))
	authRoutes.GET("/api/files", files.Index())
	authRoutes.POST("/api/files", files.Store())

	a.Routes.GET("/files/:file_id", files.Show())
}