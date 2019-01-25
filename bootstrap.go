package main

import (
	"fmt"
	"github.com/dotuancd/ezserve/app/validation"
	"gopkg.in/go-playground/validator.v8"

	. "github.com/dotuancd/ezserve/app"
	"github.com/dotuancd/ezserve/app/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

func bootApp(a *App) {
	loadConfig(a)
	initRoutes(a)
	initDatabase(a)
	migrateDatabase(a)
	registerValidators(a)
	loadViews(a)
}

func registerValidators(a *App) *App {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("required_without", validation.RequiredWithout)
	}

	return a
}

func loadConfig(a *App) *App {
	v := viper.GetViper()
	v.SetConfigType("json")
	v.SetConfigName("config")
	v.AddConfigPath("config")

	err := v.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("[ERROR] Cannot reading the config %s \n", err))
	}

	a.Config = v
	return a
}

func initRoutes(a *App) {
	a.Routes = gin.Default()
	registerGlobalMiddleware(a)
	registerRoutes(a)
}

func initDatabase(a *App) *App {
	dbConf := a.Config.GetStringMapString("database")

	c := mysql.NewConfig()
	c.DBName = dbConf["database"]
	c.Passwd = dbConf["password"]
	c.User = dbConf["user"]
	c.ParseTime = true

	dsn := c.FormatDSN()
	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic(fmt.Errorf("[ERROR] Cannot connect to database %s", err))
	}

	a.DB = db

	return a
}

func migrateDatabase(a *App) *App {
	a.DB.AutoMigrate(&models.User{})
	a.DB.AutoMigrate(&models.File{})
	return a
}

func loadViews(a *App) *App {
	a.Routes.LoadHTMLGlob("resources/views/*")
	return a
}

