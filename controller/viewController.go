package controller

import "github.com/gofiber/fiber/v2"

func Index(ctx *fiber.Ctx) error {
	return ctx.Render("index", fiber.Map{})
}

func Login(ctx *fiber.Ctx) error {
	return ctx.Render("login", fiber.Map{})
}

func Archive(ctx *fiber.Ctx) error {
	return ctx.Render("archive", fiber.Map{})
}
func Article(ctx *fiber.Ctx) error {
	return ctx.Render("article", fiber.Map{})
}
func ForgetPwd(ctx *fiber.Ctx) error {
	return ctx.Render("forget-pwd", fiber.Map{})
}
func Category(ctx *fiber.Ctx) error {
	return ctx.Render("category", fiber.Map{})
}
func Link(ctx *fiber.Ctx) error {
	return ctx.Render("link", fiber.Map{})
}
func Register(ctx *fiber.Ctx) error {
	return ctx.Render("register", fiber.Map{})
}

func Search(ctx *fiber.Ctx) error {
	return ctx.Render("search", fiber.Map{})
}

func Tag(ctx *fiber.Ctx) error {
	return ctx.Render("tag", fiber.Map{})
}
