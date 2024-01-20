package api

import (
	"github.com/gofiber/fiber/v2"
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/api/rest/handlers"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()
	rh := &rest.RestHandler{
		App: app,
	}
	setupRoutes(rh)
	app.Listen(config.ServerPort)

}

func setupRoutes(rh *rest.RestHandler) {
	// user handlers
	handlers.SetupUserRoutes(rh) //rh input because this func req this arguments

	//transactions
	//catalog
}
