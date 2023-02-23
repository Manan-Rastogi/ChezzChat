package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setUpRoutes(router *gin.Engine){
	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ALL OK",
		})
	})
	
	api :=router.Group("api/v1")

	api.POST("/", nil)
}