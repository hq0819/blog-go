package main

import (
	"blog/controller"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/middleware/skip"
	"github.com/gofiber/template/html"
	"log"
)

func main() {
	engine := html.New("template", ".html")
	engine.Delims("${{", "}}")
	app := fiber.New(fiber.Config{Views: engine, ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.Render("error", nil)
	}})
	app.Static("/static", "/static")
	handler := recover.New()
	app.Use(handler)
	//登录拦截
	app.Use(skip.New(func(ctx *fiber.Ctx) error {
		store := session.New()
		session, _ := store.Get(ctx)
		u := session.Get("user")
		fmt.Println(u)
		println(session.ID())
		if u == nil {
			return ctx.Render("login", nil)
		}
		return ctx.Next()
	}, func(c *fiber.Ctx) bool {
		fmt.Println(c.Path() == `/user/login`)
		if c.Path() == `/user/login` {
			return true
		}
		return false
	}))
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

	//业务
	group := app.Group("/user")
	group.Post("/login", controller.GetLogin)

	log.Println(app.Listen(":8002"))

}
