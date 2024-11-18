package initialize

import (
	"github.com/gofiber/fiber/v2"

	router2 "github.com/augustus281/downloader/internal/router"
)

func InitRouter() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "Downloader with Google Driver",
	})

	router := router2.NewRouter()
	router.RegisterRoutes(app)

	return app
}
