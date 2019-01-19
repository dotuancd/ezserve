package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type App struct {
	DB     *gorm.DB
	Config *viper.Viper
	Routes *gin.Engine
}

var defaultApp = &App{}

func DefaultApp() *App {
	return defaultApp
}
