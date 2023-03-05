package controllers

import (
	"net/http"
	"strconv"

	"github.com/Manan-Rastogi/chezz/helpers"
	"github.com/Manan-Rastogi/chezz/models"
	"github.com/Manan-Rastogi/chezz/services"
	"github.com/gin-gonic/gin"
)

type Controllers interface {
	CreateUserController(ctx *gin.Context)
	GetUserController(ctx *gin.Context)
	CreateGroupController(ctx *gin.Context)
	GetGroupController(ctx *gin.Context)
	CreateGroupMembersController(ctx *gin.Context)
	GetGroupMembersController(ctx *gin.Context)
	CreateGroupChatController(ctx *gin.Context)
	GetGroupChatController(ctx *gin.Context)
	CreateFriendsController(ctx *gin.Context)
	GetFriendsController(ctx *gin.Context)
}

type controllers struct {
	service services.Services
}

func NewController(s services.Services) Controllers {
	return &controllers{
		service: s,
	}
}

func (c *controllers) CreateUserController(ctx *gin.Context) {
	var user models.Users
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "Invalid Input",
		})
		return
	}

	code, msg, err := c.service.CreateUserService(user)

	ctx.JSON(code, gin.H{
		"error": err.Error(),
		"msg":   msg,
	})
}

func (c *controllers) GetUserController(ctx *gin.Context) {
	username := ctx.Param("name")

	if username == `` {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": "user request parameter cannot be blank",
			"msg": "Invalid Input",
		})
		return
	}

	code, user, err := c.service.GetUserService(username)

	ctx.JSON(code, gin.H{
		"err": err.Error(),
		"msg": user,
	})
}

func (c *controllers) CreateGroupController(ctx *gin.Context) {
	var grp models.Groups
	err := ctx.ShouldBindJSON(&grp)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "Invalid Input",
		})
		return
	}

	code, msg, err := c.service.CreateGroupService(grp)

	ctx.JSON(code, gin.H{
		"error": err.Error(),
		"msg":   msg,
	})
}

func (c *controllers) GetGroupController(ctx *gin.Context) {
	groupname := ctx.Param("name")

	if groupname == `` {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": "name request parameter cannot be blank",
			"msg": "Invalid Input",
		})
		return
	}

	code, grp, err := c.service.GetGroupService(groupname)

	ctx.JSON(code, gin.H{
		"err": err.Error(),
		"msg": grp,
	})
}

func (c *controllers) CreateGroupMembersController(ctx *gin.Context) {
	var grpMems models.GroupMembers
	err := ctx.ShouldBindJSON(&grpMems)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "Invalid Input",
		})
		return
	}

	code, msg, err := c.service.CreateGroupMembersService(grpMems)

	ctx.JSON(code, gin.H{
		"error": err.Error(),
		"msg":   msg,
	})
}

func (c *controllers) GetGroupMembersController(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == `` {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": "id request parameter cannot be blank",
			"msg": "Invalid Input",
		})
		return
	}
	Id, err := strconv.Atoi(id)
	if err != nil{
		helpers.Logger.Errorf("Error in converting id to int: %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": "id request parameter must be a number",
			"msg": "Invalid Input",
		})
		return
	}

	code, grpMems, err := c.service.GetGroupMembersService(Id)

	ctx.JSON(code, gin.H{
		"err": err.Error(),
		"msg": grpMems,
	})
}

func (c *controllers) CreateGroupChatController(ctx *gin.Context) {
	var grpChats models.GroupChats
	err := ctx.ShouldBindJSON(&grpChats)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "Invalid Input",
		})
		return
	}

	code, msg, err := c.service.CreateGroupChatService(grpChats)

	ctx.JSON(code, gin.H{
		"error": err.Error(),
		"msg":   msg,
	})
}

func (c *controllers) GetGroupChatController(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == `` {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": "id request parameter cannot be blank",
			"msg": "Invalid Input",
		})
		return
	}
	Id, err := strconv.Atoi(id)
	if err != nil{
		helpers.Logger.Errorf("Error in converting id to int: %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": "id request parameter must be a number",
			"msg": "Invalid Input",
		})
		return
	}

	code, grpChats, err := c.service.GetGroupChatService(Id)

	ctx.JSON(code, gin.H{
		"err": err.Error(),
		"msg": grpChats,
	})
}

func (c *controllers) CreateFriendsController(ctx *gin.Context) {
	var friends models.Friends
	err := ctx.ShouldBindJSON(&friends)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "Invalid Input",
		})
		return
	}

	code, msg, err := c.service.CreateFriendsService(friends)

	ctx.JSON(code, gin.H{
		"error": err.Error(),
		"msg":   msg,
	})
}

func (c *controllers) GetFriendsController(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == `` {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": "id request parameter cannot be blank",
			"msg": "Invalid Input",
		})
		return
	}
	Id, err := strconv.Atoi(id)
	if err != nil{
		helpers.Logger.Errorf("Error in converting id to int: %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": "id request parameter must be a number",
			"msg": "Invalid Input",
		})
		return
	}

	code, friends, err := c.service.GetFriendsService(Id)

	ctx.JSON(code, gin.H{
		"err": err.Error(),
		"msg": friends,
	})
}
