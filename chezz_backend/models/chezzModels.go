package models

type Users struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name" binding:"required,min=5,max=30"`
	Email     string `json:"email" db:"email" binding:"required,email,endswith=@chezz.com,min=15,max=30"`
	Password  string `json:"password" db:"password" binding:"required,min=8,max=30"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

type Groups struct {
	Id        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Creator   int    `json:"creator" db:"creator"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

type GroupMembers struct {
	Id       int    `json:"id" db:"id"`
	GroupId  int    `json:"group_id" db:"groupId"`
	UserId   int    `json:"user_id" db:"userId"`
	JoinedAt string `json:"joined_at" db:"joined_at"`
}

type GroupChats struct {
	Id        int    `json:"id" db:"id"`
	GroupId   int    `json:"group_id" db:"groupId"`
	UserId    int    `json:"user_id" db:"userId"`
	Message   string `json:"message" db:"message"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

type Friends struct {
	Id        int    `json:"id" db:"id"`
	UserId    int    `json:"user_id" db:"userId"`
	FriendId  int    `json:"friend_id" db:"friendId"`
	CreatedAt string `json:"created_at" db:"created_at"`
}
