package main

import (
	"io"
	"os"
	"regexp"

	"github.com/Manan-Rastogi/chezz/configs"
	"github.com/Manan-Rastogi/chezz/helpers"
	"github.com/Manan-Rastogi/chezz/middleware"
	"github.com/Manan-Rastogi/chezz/models"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

var rxURL = regexp.MustCompile(`^/regexp\d*`)

func main() {
	helpers.Logger.Info("Chezz Service Started!!")
	models.DB = models.Init()

	logFile := generateGinLoggerFile()
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	router := gin.New()
	router.Use(gin.Recovery(), middleware.Logger(logFile), cors.New(middleware.CorsPolicy()))
	setUpRoutes(router)

	router.Run(configs.CONFIG.PORT)
}
