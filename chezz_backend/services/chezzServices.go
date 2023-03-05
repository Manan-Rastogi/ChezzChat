package services

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Manan-Rastogi/chezz/models"
)

type Services interface {
	CreateUserService(user models.Users) (code int, msg string, err error)
	GetUserService(username string) (code int, user models.Users, err error)
	CreateGroupService(grp models.Groups) (code int, msg string, err error)
	GetGroupService(name string) (code int, grp models.Groups, err error)
	CreateGroupMembersService(gm models.GroupMembers) (code int, msg string, err error)
	GetGroupMembersService(id int) (code int, gm []models.GroupMembers, err error)
	CreateGroupChatService(gc models.GroupChats) (code int, msg string, err error)
	GetGroupChatService(id int) (code int, gc []models.GroupChats, err error)
	CreateFriendsService(friend models.Friends) (code int, msg string, err error)
	GetFriendsService(id int) (code int, friends []models.Friends, err error)
}

type services struct {
}

func NewService() Services {
	return &services{}
}

func (s *services) CreateUserService(user models.Users) (code int, msg string, err error) {
	err = models.CreateNewUser(models.DB, user)
	if err != nil {
		return http.StatusInternalServerError, "Failed to create new user", err
	}

	return http.StatusOK, "User Created", fmt.Errorf("")
}

func (s *services) GetUserService(username string) (code int, user models.Users, err error) {
	user, err = models.GetUserByName(models.DB, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return http.StatusBadRequest, user, err
		}
		return http.StatusInternalServerError, user, err
	}
	user.Password = `*********`
	return http.StatusOK, user, fmt.Errorf("")
}

func (s *services) CreateGroupService(grp models.Groups) (code int, msg string, err error) {
	err = models.CreateNewGroup(models.DB, grp)
	if err != nil {
		return http.StatusInternalServerError, "Failed to create new group", err
	}

	return http.StatusOK, "Group Created", fmt.Errorf("")
}

func (s *services) GetGroupService(name string) (code int, grp models.Groups, err error) {
	grp, err = models.GetGroupByName(models.DB, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return http.StatusBadRequest, grp, err
		}
		return http.StatusInternalServerError, grp, err
	}
	
	return http.StatusOK, grp, fmt.Errorf("")
}

func (s *services) CreateGroupMembersService(gm models.GroupMembers) (code int, msg string, err error) {
	err = models.CreateGroupMembers(models.DB, gm)
	if err != nil {
		return http.StatusInternalServerError, "Failed to create new group members", err
	}

	return http.StatusOK, "Group Member Created", fmt.Errorf("")
}

func (s *services) GetGroupMembersService(id int) (code int, gm []models.GroupMembers, err error) {
	gm, err = models.GetGroupMembersByGroupId(models.DB, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return http.StatusBadRequest, gm, err
		}
		return http.StatusInternalServerError, gm, err
	}
	
	return http.StatusOK, gm, fmt.Errorf("")
}

func (s *services) CreateGroupChatService(gc models.GroupChats) (code int, msg string, err error) {
	err = models.CreateGroupChat(models.DB, gc)
	if err != nil {
		return http.StatusInternalServerError, "Failed to create new group Chat", err
	}

	return http.StatusOK, "Group Chat Created", fmt.Errorf("")
}

func (s *services) GetGroupChatService(id int) (code int, gc []models.GroupChats, err error) {
	gc, err = models.GetGroupChatsByGroupId(models.DB, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return http.StatusBadRequest, gc, err
		}
		return http.StatusInternalServerError, gc, err
	}
	
	return http.StatusOK, gc, fmt.Errorf("")
}

func (s *services) CreateFriendsService(friend models.Friends) (code int, msg string, err error) {
	err = models.CreateFriends(models.DB, friend)
	if err != nil {
		return http.StatusInternalServerError, "Failed to Add new friend", err
	}

	return http.StatusOK, "Friend Added", fmt.Errorf("")
}

func (s *services) GetFriendsService(id int) (code int, friends []models.Friends, err error) {
	friends, err = models.GetFriendsByUserId(models.DB, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return http.StatusBadRequest, friends, err
		}
		return http.StatusInternalServerError, friends, err
	}
	
	return http.StatusOK, friends, fmt.Errorf("")
}
