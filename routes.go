package main

import (
	"github.com/dotuancd/ezserve/app"
	"github.com/dotuancd/ezserve/app/http/controllers"
	"github.com/dotuancd/ezserve/app/http/errors"
	"github.com/dotuancd/ezserve/app/http/middlewares"
)

func registerRoutes(a *app.App) {
	registerHomeRoutes(a)
	registerFileRoutes(a)
	registerAssetsRoutes(a)
}

func registerHomeRoutes(a *app.App) {
	home := controllers.HomeController{
		App:a,
	}

	a.Routes.GET("/", errors.Handler(home.Index))
}

func registerFileRoutes(a *app.App) {
	files := &controllers.FileHandler{
		App: a,
	}

	r := a.Routes
	r.GET("/files/:file_id/:filename", errors.Handler(files.Show()))

	authRoutes := r.Group("/", middlewares.UserAuth(a))
	authRoutes.GET("/api/files", errors.Handler(files.Index()))
	authRoutes.POST("/api/files", errors.Handler(files.Store()))
}

func registerAssetsRoutes(a *app.App) {
	a.Routes.Static("/assets/", "public")
	a.Routes.StaticFile("/favicon.ico", "public/favicon.ico")
}