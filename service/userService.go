package service

import (
	"blog/configuration"
	"blog/model"
	"gorm.io/gorm"
)

var Db = configuration.GetDbInstance()

func UserLogin(u *model.User) *model.User {
	m := &model.User{}
	Db.Where(u).First(m)
	return m
}

func AddUser(u *model.User) {

	_ = Db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(u).Error
	})
}
