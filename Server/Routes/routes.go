package routes

import (
	controller "server/Controller"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {

	app.Get("/blogdetail/:id", controller.GetBlogDetail)
	app.Get("/bloglist", controller.GetBlogList)
	app.Post("/createblog", controller.CreateBlog)
	app.Put("/updateblog/:id", controller.UpdateBlog)
	app.Delete("/deleteblog/:id", controller.DeleteBlog)
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Get("/api/user", controller.User)
	app.Post("/api/logout", controller.Logout)
}
