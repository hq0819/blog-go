package controller

import (
	"blog/model"
	"blog/service"
	"blog/util"
	"encoding/gob"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"time"
)

func GetLogin(ctx *fiber.Ctx) error {
	validate := validator.New()
	store := session.New(session.Config{Storage: util.GetInstance()})
	get, _ := store.Get(ctx)
	va := &model.User{0, ctx.FormValue("username"), ctx.FormValue("passwd")}
	err2 := validate.Struct(va)
	if err2 != nil {
		return ctx.Render("newLogin", nil)
	}

	login := service.UserLogin(va)

	if login.Username == "" {
		return ctx.Render("newLogin", nil)
	}
	gob.Register(model.User{})
	get.SetExpiry(time.Duration(60*5) * time.Second)
	get.Set("user", login)
	err := get.Save()
	if err != nil {
		fmt.Println(err)
	}
	return ctx.Render("index", *login)
}
