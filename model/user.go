package model

type User struct {
	UserId   int32  `json:"userId"`
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
}

func NewUser(username string, passwd string) *User {
	return &User{Username: username, Passwd: passwd}
}
