package model

type User struct {
	UserId   int32  `json:"userId" gorm:"primarykey;auto_increment;column:userId;type:int(20)" `
	Username string `json:"username" gorm:"column:username;type:varchar(20)" validate:"required"`
	Passwd   string `json:"passwd" gorm:"column:passwd;type:varchar(20)" validate:"required"`
}

func NewUser(username string, passwd string) *User {
	return &User{Username: username, Passwd: passwd}
}

func (User) TableName() string {
	return "tb_user"
}
