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

	_, err = db.Exec(insertQuery, user.Name, user.Email, user.Password)
	if err != nil{
		helpers.Logger.Errorf("Error Creating user: %v", err.Error())
		return
	}

	return
}


func CreateNewGroup(db *sqlx.DB, group Groups) (err error) {
	insertQuery := "INSERT into `groups` (name, creator) VALUES(?, ?);"

	// TODO - A user can be a part of maximum 10 groups

	_, err = db.Exec(insertQuery, group.Name, group.Creator)
	if err != nil{
		helpers.Logger.Errorf("Error Creating group: %v", err.Error())
		return
	}

	return 
}


func CreateGroupMembers(db *sqlx.DB, gm GroupMembers) (err error) {
	insertQuery := `INSERT INTO group_members(groupId, userId) VALUES(?, ?);`

	_, err = db.Exec(insertQuery, gm.GroupId, gm.UserId)
	if err != nil{
		helpers.Logger.Errorf("Error Creating group members: %v", err.Error())
		return
	}

	return 
}

func CreateGroupChat(db *sqlx.DB, gc GroupChats) (err error) {
	insertQuery := `INSERT INTO group_chat (groupId, userId, message) VALUES(?, ?, ?);`

	// TODO - Limit the size of Msg.

	_, err = db.Exec(insertQuery, gc.GroupId, gc.UserId, gc.Message)
	if err != nil{
		helpers.Logger.Errorf("Error Creating group chat: %v", err.Error())
		return
	}

	return 
}

func CreateFriends(db *sqlx.DB, friends Friends) (err error) {
	insertQuery := `INSERT INTO friends (userId, friendId) VALUES(?, ?), (?, ?);`

	// TODO - Check if two are already friends

	_, err = db.Exec(insertQuery, friends.UserId, friends.FriendId, friends.FriendId, friends.UserId)
	if err != nil{
		helpers.Logger.Errorf("Error Creating friends: %v", err.Error())
		return
	}

	return 
}

