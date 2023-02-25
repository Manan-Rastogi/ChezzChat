package models

import "time"

type Users struct {
	Id        int       `json:"id" db:"-"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Groups struct {
	Id        int       `json:"id" db:"-"`
	Name      string    `json:"name" db:"name"`
	Creator   string    `json:"creator" db:"creator"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type GroupMembers struct {
	Id       int       `json:"id" db:"-"`
	GroupId  string    `json:"group_id" db:"groupId"`
	UserId   int       `json:"user_id" db:"userId"`
	JoinedAt time.Time `json:"joined_at" db:"joined_at"`
}

type GroupChats struct {
	Id        int       `json:"id" db:"-"`
	GroupId   string    `json:"group_id" db:"groupId"`
	UserId    int       `json:"user_id" db:"userId"`
	Message   string    `json:"message" db:"message"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Friends struct {
	Id        int       `json:"id" db:"-"`
	UserId    int       `json:"user_id" db:"userId"`
	FriendId  int       `json:"friend_id" db:"friendId"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
