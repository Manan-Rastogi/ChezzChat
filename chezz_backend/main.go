package main

import (
	"context"
	"regexp"

	"time"

	"github.com/Manan-Rastogi/chezz/configs"
	"github.com/Manan-Rastogi/chezz/helpers"
	"github.com/Manan-Rastogi/chezz/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/logger"
)

var rxURL = regexp.MustCompile(`^/regexp\d*`)

func main() {
	helpers.Logger.Info("Chezz Service Started!!")
	start := time.Now()
	ctx := context.Background()

	db := &models.DB{}
	db.ConfigureDB(ctx)

	ctx = context.WithValue(ctx, "start", start)

	router := gin.New()
	router.Use(gin.Recovery(), logger.SetLogger(
		
		logger.WithUTC(true),
		logger.WithSkipPathRegexp(rxURL),
	), cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "UPDATE", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		// ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
	}))
	setUpRoutes(router)

	router.Run(configs.CONFIG.PORT)
}
