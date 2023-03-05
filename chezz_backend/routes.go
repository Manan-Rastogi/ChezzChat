package main

import (
	"net/http"

	"github.com/Manan-Rastogi/chezz/controllers"
	"github.com/Manan-Rastogi/chezz/services"
	"github.com/gin-gonic/gin"
)

var (
	service    services.Services       = services.NewService()
	controller controllers.Controllers = controllers.NewController(service)
)

func setUpRoutes(router *gin.Engine) {

	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "ALL OK"})
	})

	api := router.Group("api/v1")

	api.POST("/createUser", controller.CreateUserController)
	api.POST("/getUser/:name", controller.GetUserController)
	api.POST("/createGroup", controller.CreateGroupController)
	api.POST("/getGroup/:name", controller.GetGroupController)
	api.POST("/createGroupMember", controller.CreateGroupMembersController)
	api.POST("/getGroupMembers/:id", controller.GetGroupMembersController)
	api.POST("/createGroupChat", controller.CreateGroupChatController)
	api.POST("/getGroupChats/:id", controller.GetGroupChatController)
	api.POST("/addFriend", controller.CreateFriendsController)
	api.POST("/getFriends/:id", controller.GetFriendsController)
}
