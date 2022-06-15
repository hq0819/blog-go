package main

import (
	"blog/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("template", ".html")
	app := fiber.New(fiber.Config{Views: engine, ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.Render("error", nil)
	}})
	app.Static("/static", "/static")
	store := session.New()
	handler := recover.New()
	app.Use(handler)
	//登录拦截
	app.Use("/", func(ctx *fiber.Ctx) error {
		session, _ := store.Get(ctx)
		u := session.Get("user")
		if u == nil {
			return ctx.Render("login", nil)
		}
		return ctx.Next()
	})
	//页面控制器
	app.Get("/index", controller.Index)
	app.Get("/loginPage", controller.Login)
	app.Get("/archive", controller.Archive)
	app.Get("/article", controller.Article)
	app.Get("/forgetPwd", controller.ForgetPwd)
	app.Get("/category", controller.Category)
	app.Get("/link", controller.Link)
	app.Get("/register", controller.Register)
	app.Get("/search", controller.Search)
	app.Get("/tag", controller.Tag)

	app.Listen(":8002")

}
