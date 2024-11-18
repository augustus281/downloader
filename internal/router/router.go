package router

import "github.com/gofiber/fiber/v2"

type Router struct{}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) InitHealthCheck(router fiber.Router) {
	router.Get("/healthcheck", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{"status": "OK"})
	})
}

func (r *Router) RegisterRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	r.InitHealthCheck(api)
}
