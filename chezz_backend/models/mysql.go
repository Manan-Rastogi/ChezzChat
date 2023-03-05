package models

import (
	"fmt"
	"log"

	// "github.com/Manan-Rastogi/chezz/configs"
	"github.com/Manan-Rastogi/chezz/helpers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Init() *sqlx.DB {
	// user := configs.CONFIG.USER
	// password := configs.CONFIG.PASSWORD
	// ip := configs.CONFIG.IP
	// dbName := configs.CONFIG.DB

	user := "root"
	password := "password"
	ip := "localhost:3306"
	dbName := "chezz"

	connString := fmt.Sprintf(`%v:%v@tcp(%v)/%v`, user, password, ip, dbName)

	var err error
	db, err := sqlx.Open("mysql", connString)
	if err != nil {
		log.Fatalf("Error establishing connection to db: %v", err.Error())
	}

	// Ping DB to check connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error verifing connection to db: %v", err.Error())
	}

	helpers.Logger.Info("DB Connection Successful!")
	return db
}

func CreateNewUser(db *sqlx.DB, user Users) (err error) {
	insertQuery := `INSERT INTO users (name, email, password) VALUES(?, ?, ?);`

	// TODO - check if user exists

	_, err = db.Exec(insertQuery, user.Name, user.Email, user.Password)
	if err != nil {
		helpers.Logger.Errorf("Error Creating user: %v", err.Error())
		return
	}

	return
}

func CreateNewGroup(db *sqlx.DB, group Groups) (err error) {
	insertQuery := "INSERT into `groups` (name, creator) VALUES(?, ?);"

	// TODO - A user can be a part of maximum 10 groups
	// TODO - check if group exists

	_, err = db.Exec(insertQuery, group.Name, group.Creator)
	if err != nil {
		helpers.Logger.Errorf("Error Creating group: %v", err.Error())
		return
	}

	return
}

func CreateGroupMembers(db *sqlx.DB, gm GroupMembers) (err error) {
	insertQuery := `INSERT INTO group_members(groupId, userId) VALUES(?, ?);`

	_, err = db.Exec(insertQuery, gm.GroupId, gm.UserId)
	if err != nil {
		helpers.Logger.Errorf("Error Creating group members: %v", err.Error())
		return
	}

	return
}

func CreateGroupChat(db *sqlx.DB, gc GroupChats) (err error) {
	insertQuery := `INSERT INTO group_chat (groupId, userId, message) VALUES(?, ?, ?);`

	// TODO - Limit the size of Msg.

	_, err = db.Exec(insertQuery, gc.GroupId, gc.UserId, gc.Message)
	if err != nil {
		helpers.Logger.Errorf("Error Creating group chat: %v", err.Error())
		return
	}

	return
}

func CreateFriends(db *sqlx.DB, friends Friends) (err error) {
	insertQuery := `INSERT INTO friends (userId, friendId) VALUES(?, ?), (?, ?);`

	// TODO - Check if two are already friends

	_, err = db.Exec(insertQuery, friends.UserId, friends.FriendId, friends.FriendId, friends.UserId)
	if err != nil {
		helpers.Logger.Errorf("Error Creating friends: %v", err.Error())
		return
	}

	return
}

func GetUserByName(db *sqlx.DB, name string) (user Users, err error) {
	selectQuery := `Select * from users where name = ?`

	err = db.Get(&user, selectQuery, name)
	if err != nil {
		helpers.Logger.Errorf("Error getting details for user %v: %v", name, err.Error())
		return
	}

	helpers.Logger.Infof("User Details fetched for %v", name)

	return
}

func GetGroupByName(db *sqlx.DB, name string) (grp Groups, err error) {
	selectQuery := "Select * from `groups` where name = ?"

	err = db.Get(&grp, selectQuery, name)
	if err != nil {
		helpers.Logger.Errorf("Error getting details for group %v: %v", name, err.Error())
		return
	}

	helpers.Logger.Infof("Group Details fetched for %v", name)
	return
}

func GetGroupMembersByGroupId(db *sqlx.DB, gId int) (gms []GroupMembers, err error) {
	selectQuery := "Select * from `group_members` where groupId = ?"

	err = db.Select(&gms, selectQuery, gId)
	if err != nil {
		helpers.Logger.Errorf("Error getting group members for group %v: %v", gId, err.Error())
		return
	}

	helpers.Logger.Infof("Group Members Details fetched for %v", gId)
	return
}

func GetGroupChatsByGroupId(db *sqlx.DB, gId int) (gms []GroupChats, err error) {
	selectQuery := "Select * from `group_chat` where groupId = ? order by created_at desc limit 100"
	// Add index for created_at
	err = db.Select(&gms, selectQuery, gId)
	if err != nil {
		helpers.Logger.Errorf("Error getting group chats for group %v: %v", gId, err.Error())
		return
	}

	ReverseSlice(gms)

	helpers.Logger.Infof("Group Chat Details fetched for %v", gId)
	return
}

func GetFriendsByUserId(db *sqlx.DB, uId int) (friends []Friends, err error) {
	selectQuery := `Select * from friends where userId = ?`

	err = db.Select(&friends, selectQuery, uId)
	if err != nil {
		helpers.Logger.Errorf("Error getting friends for user %v: %v", uId, err.Error())
		return
	}

	helpers.Logger.Infof("User Details fetched for %v", uId)
	return
}
