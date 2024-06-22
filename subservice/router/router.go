package router

import (
	"github.com/daifiyum/cat-box/subservice/handler"
	"github.com/daifiyum/cat-box/subservice/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// index
	app.Static("/", "./resources/ui/sub")
	// hello
	api := app.Group("/api")
	api.Get("/", handler.Hello)
	// Auth
	api.Post("/login", handler.Login)

	// User
	user := api.Group("/user")
	user.Get("/", handler.IsRegistered)
	user.Post("/", handler.CreateUser)

	// Option
	option := api.Group("/option")
	option.Post("/", middleware.Protected(), handler.UpdateOption)
	option.Get("/", middleware.Protected(), handler.GetOption)

	// Subscribe
	subscribe := api.Group("/subscribe")
	subscribe.Get("/", middleware.Protected(), handler.GetAllSubscribe)
	subscribe.Post("/", middleware.Protected(), handler.CreateSubscribe)
	subscribe.Delete("/:id", middleware.Protected(), handler.DeleteSubscribe)
	subscribe.Put("/:id", middleware.Protected(), handler.EditSubscribe)
	subscribe.Patch("/:id", middleware.Protected(), handler.OperateSubscribe)

	// Singbox
	singbox := api.Group("/singbox")
	singbox.Put("/", middleware.Protected(), handler.StopSingbox)
}
