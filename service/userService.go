package service

import (
	"blog/model"
)

func UserLogin(u *model.User) *model.User {
	return model.NewUser("heqin", "123")
}
